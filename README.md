#batrium-udp2http-bridge

This service should Listen for UDP broadcasts from the Batrium Watchmon 4+ and
collect all the metrics and expose them on as JSON on individual http routes.

## BUILD
docker build -t batrium-udp2http-bridge . ;

## RUN
docker run -it batrium-udp2http-bridge:latest ;
