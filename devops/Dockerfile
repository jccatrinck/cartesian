FROM golang:1.15.7-alpine3.13 as build

ARG PACKAGE=""
ARG WORKPATH="${GOPATH}/src/${PACKAGE}"

RUN apk add --no-cache git

WORKDIR $WORKPATH

#                                                                                  .+                               
#                                                                                  oh                               
#                                                                                  sd                               
#                                                      `y`                         yd                               
#                    ..                                /Mo                         ym`                              
#                  `+`h:                               dNN-                        hm`                              
#                  `d`sy                              -NmNh                        hN`                              
#                   /od/                              yNmmN/                       hN.                              
#                    .m.                             .NmmmNd`                      dM.                              
#                     s+                             sNmmmmNo                      dM.                              
#                     -h                          `./Nmmmmmmmo/:--.``              dM.                              
#                      h-                  `.-:+sydmmmmmmmmmmmNNmmmdhhy+.        --dM/-.                            
#                      +s                :oydmNNmmmmmmmmmmmmmmmmNmmdhs/.        `/smNmo:                            
#                      -m.`               `-:/oshddmmmmmmmmmmNho/:-`          `.``+mmm`                             
#                     -dmdh/.//:-.`               `hNmmmNNmmmNy           `./ohm:dNmmd`                             
#                     :NmmmN+mNNmmhs+-.`      ````.Nmmd/sssmmNs``     `./shmNmmN+yood-                              
#                      +dmds-+NmmmmmNNmhyo++oshddddmmm+-ho-yNmmdhysooydmNmmmmmmNo  -d`                              
#                       `m+  `mNmmmmmmmmmmmmmmmmmmmmNy`    -mmmmmmmmmmmmmmmmmmmNy                                   
#                        sh   +Nmmmmmmmmmmmmmmmmmmmmd-     `yNmmmmmmmmmmmmmmmmmNy                                   
#                        :N.  `Nmmmmmmmmmmmmmmmmmmmmm.      yNmNhyhmmmmmmmmmmmmNh                                   
#                        `m/   dNmmmmmmmmmmdyo+oNmmNh      `+NmN:  `-+sdmmmmmmmmm                                   
#                         yh   hNmmmmmho/-`     mNmNoo     -NmmN-       `-/oydmmN`                                  
#                         /N` `sso/:.           dNmmNy     `NmmN:              ``                                   
#                         .N/                   mmmmNd``  /-mmmNs                                                   
#                          hy                  .Nmmmmmm+. mNmmmNd                                                   
#                          +N`                 /NmmmmmmdyyNmmmmmN-                                                  
#                          .M:                 yNmmmmmmmmmmmmmmmNy                                                  
#                           ds                `NmmmmmmmmmmmmmmmmmN.                                                 
#                           om                /NmmmmmmmmmmmmmmmmmNy                                                 
#                           -M-               dNmmmmmmmmmmmmmmmmmmN-                                                
#                            do              :NmmmmmmmmmmmmmmmmmmmNh                                                
#                            od              hNmmmmmmmmmmmmmmmmmmmmN/                                               
#                            -M.            -Nmmmmmmmmmmmmmmmmmmmmmmm.                                              
#                            `m+            hNmmmmmmmmmmmmmmmmmmmmmmNy                                              
#                             sh           -NmmmmmmmmmmmmmmmmmmmmmmmmN/                                             
#                             :M`          hNmmmmmmmmmmmmmmmmmmmmmmmmmm.                                            
#                             `m/         :NmmmmmmmmmmmmmmmmmmmmmmmmmmNh`                                           
#                              sy        `dNmmmmmmmmmmmmmmmmmmmmmmmmmmmNs                                           
#                              :N`       /NmmmmmmmmmmmmmmmmmmmmmmmmmmmmmN/                                          
#                               :`      `mNNNNNNNNNNNNNNNNNmmmmmNNNNmmmmNm-                                         
#                                       -++/////syyys///++++++syyyyoooooss+                                         
#                                              -hhhhh`       .hdddd- 
# Advice
# Every line below will not been cached if you pass: --build-arg CACHE_BUST=$(date +%s) api
ARG CACHE_BUST=0

COPY . .

# Build the API executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags "-X 'main.Version=$(date '+%Y%m%d%H%M%S')'" -o /root/app *.go

# Lightweight reliable docker image for running the API
FROM alpine:3.13

WORKDIR /root/

# Copy API executable
COPY --from=build /root/app .

ENTRYPOINT ./app