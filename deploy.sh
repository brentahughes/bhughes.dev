APP_NAME=personal-website

RUN_FLAG="-d --restart=always"
if [ "$1" == "debug" ]; then
    RUN_FLAG="--rm"
fi

echo "Building $APP_NAME image"
docker build --no-cache -t $APP_NAME .

echo "Removing $APP_NAME container if it exists"
docker rm -f $APP_NAME

echo "Running $APP_NAME container"
docker run $RUN_FLAG --network traefik --name $APP_NAME \
    -l "traefik.http.routers.personalsite.rule=Host(\`brentahughes.com\`, \`www.brentahughes.com\`, \`bhughes.dev \`, \`www.bhughes.dev\`)" \
    -l "traefik.http.routers.personalsite.entrypoints=secure-web" \
    -l "traefik.http.routers.personalsite.tls=true" \
    -l "traefik.http.routers.personalsite.tls.certresolver=letsencrypt" \
    -l "traefik.http.routers.personalsite.tls.domains[0].main=brentahughes.com" \
    -l "traefik.http.routers.personalsite.tls.domains[0].sans=*.brentahughes.com" \
    -l "traefik.http.routers.personalsite.tls.domains[1].main=bhughes.dev" \
    -l "traefik.http.routers.personalsite.tls.domains[1].sans=*.bhughes.dev" \
    -l "traefik.enable=true" \
    -p 8800:8080 \
    $APP_NAME
