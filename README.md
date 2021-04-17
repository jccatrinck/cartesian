# Cartesian API

Cartesian API for spatial data processing using geometry techniques to search points in 2d plane.

---
### Summary

- [Quick-start](#quick-start)
- [Structure](#storage)
- [Techniques](#techniques)
- [Storage](#storage)
- [Testing](#testing)
- [Handlers](#handlers)
- [Future work](#future-work)
---
### Quick-start

The only configuration needed is to create the `.env` file.

There is a `.env-example` that is ready to use, you can just rename it.

Use command below to run the API:

```shell
> make run
> make logs
```

To see all commands available:

```shell
> make help
```
---
### Structure
```
├── data          // API data
├── devops        // Docker related files
├── handlers      // API endpoints
├── libs          // Utilities with no business logic
├── middlwares    // Custom middlewares
├── services      // Business logic as services
├── storage       // Storage interface and implementations
├── .env-example  // Template to create .env file
├── api.go        // Main entrypoint
└── Makefile      // Commands build, run ...
```
---
### Techniques

<table>
  <tr>
    <td rowspan="3"> 
      <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/de/TaxicabGeometryCircle.svg/100px-TaxicabGeometryCircle.svg.png" />
    </td>
    <td>
      <i>Blue: searched point</i>
    </td>
  </tr>
  <tr>
    <td>
      <i>Red: searched distance</i>
    </td>
  </tr>
  <tr>
    <td>
      <i>Square: all points inside a spatial index</i>
    </td>
  </tr>
</table>

Taxicab circle is a set of points that forms a square with a fixed Manhattan distance.

Using MySQL MBR (minimum bounding rectangles) functions we are capable to take advantage of spatial indexes.

```sql
SELECT x, y
FROM point, (
  SELECT
    # Taxicab circle (in red in the image)
    POLYGON(
      LINESTRING(
        POINT(x, y+d),
        POINT(x+d, y),
        POINT(x, y-d),
        POINT(x-d, y),
        POINT(x, y+d)
      )
    ) AS circle
	FROM (
		SELECT
			:x AS x,
			:y AS y,
			:distance +1e-9 AS d
	) params
) taxicab
# Filter entire square using index
WHERE MBRCONTAINS(taxicab.circle, point)
# Filter only point inside Manhattan/taxicab circle
AND ST_CONTAINS(taxicab.circle, point)
```
Firstly filter points inside square using index later filter only inside Manhattan distance.

According to MySQL documentation if there is too many rows inside the filtered data subset the index will not be used, so there is a fallback:

```sql
SELECT a.x, a.y
FROM point a, (
		SELECT
			:x AS x,
			:y AS y,
			:distance +1e-9 AS d
) b
# Manhattan distance: |x1-x2| + |y1-y2|
WHERE ABS(a.x-b.x)+ABS(a.y-b.y) <= d
# Use primary key (x, y)
AND a.x BETWEEN b.x-d AND b.x+d
AND a.y BETWEEN b.y-d AND b.y+d
```
---
### Storage

The API was designed to accept any kind of storage system.

This is configurable in `.env` file using `API_STORAGE_TYPE`.

There are two storage types implemented:
- **Memory**: Small datasets.
- **MySQL**: Large datasets.

Services define their needs through interfaces and sub-packages of storage package will be responsible to implement these definitions.

```go
// services/points/storage.go
package points

// [...]

type Storage interface {
	LoadPoints(reader io.ReadSeeker) error
	GetPointsByDistance(point model.Point, distance int) ([]model.RelativePoint, error)
}

// storage/storage.go
package storage

import (
	// [...]
	"github.com/jccatrinck/cartesian/services/points"
	// [...]
)

type Storage interface {
	points.Storage
}
```
---
### Testing

In testing stage we create a entire environment using Docker compose to test everything.

There is no configuration needed and can be done using command below:

```shell
> make test
```
---
### Handlers
Accept and respond with JSON.

#### Authentication Header

`Authorization: Bearer ${API_KEY}`

*If `API_KEY` is empty authentication will be disabled

#### Response codes
##### 200 - OK: The request has succeeded.
##### 204 - No Content: The request has succeeded and there is no data
##### 404 - Not Found: The server can not find the requested resource.
##### 500 - Internal Server Error: The server has encountered a unhandled error


#### List points

List points inside a distance

**Endpoint** : `GET /points?distance={number}&x={number}&y={number}` 

**Payload** : `N/A`

**Response** : `[]RelativePoints`

**Example response:**
```json
[
   {
      "x": 84,
      "y": -34,
      "distance": 1
   }
]
```
---
### Future work

Use redis in memory storage type.

Deploy to the cloud using Terraform within the CD pipeline with Google App Engine and Google Cloud SQL.

Build a React front-end interface to interact with API.