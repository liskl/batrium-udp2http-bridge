#batrium-udp2http-bridge

This Package should Listen for UDP broadcasts from the Batrium Watchmon 4+ and
collect all the metrics to allow exposing that same data as JSON on individual
routes.

## BUILD
docker build -t batrium-udp2http-bridge . ;

## RUN
docker run -it batrium-udp2http-bridge:latest ;
