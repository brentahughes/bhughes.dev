APP_NAME=personal-website
VERSION=v1

echo "Building $APP_NAME image"
docker build --platform linux/amd64 -t docker.brentahughes.com:443/${APP_NAME}:${VERSION} .
docker push docker.brentahughes.com:443/${APP_NAME}:${VERSION}

echo "Running $APP_NAME"
kubectl apply -f resources/
