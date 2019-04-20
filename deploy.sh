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
docker run $RUN_FLAG --name $APP_NAME \
    -p 8800:8080 \
    $APP_NAME
