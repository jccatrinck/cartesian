package statements

// LoadPoints into database
const LoadPoints = `
INSERT INTO point (x, y) VALUES (:x, :y)
`

// PointsByDistancePK uses primary key indexing to filter relative points by distance
const PointsByDistancePK = `
SELECT
	a.x,
	a.y,
	ABS(a.x-b.x)+ABS(a.y-b.y) distance
FROM point a, (
	SELECT
		:x x,
		:y y,
		:distance d
) b
WHERE ABS(a.x-b.x)+ABS(a.y-b.y) <= d
AND a.x BETWEEN b.x-d AND b.x+d
AND a.y BETWEEN b.y-d AND b.y+d
ORDER BY distance
`

// PointsByDistanceSpatial uses spatial indexing to filter relative points by distance
const PointsByDistanceSpatial = `
SELECT
	a.x,
	a.y,
	ABS(a.x-b.x)+ABS(a.y-b.y) distance
FROM point a, (
	SELECT
		x, y, d, taxicab_circle(x, y, d) manhattan
	FROM (
		SELECT
			:x x,
			:y y,
			:distance d
	) params
) b
WHERE MBRCONTAINS(manhattan, point)
AND ST_CONTAINS(manhattan, point)
ORDER BY distance
`
