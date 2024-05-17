ROOT_DIR=`pwd`
OUT_DIR=bin

if [ ! -d ${ROOT_DIR}/${OUT_DIR} ]; then
  mkdir ${ROOT_DIR}/${OUT_DIR}
fi

GOOS=linux 
GOARCH=amd64

echo "build" ${ROOT_DIR}/cmd/api "..."
cd ${ROOT_DIR}/cmd/api && GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${ROOT_DIR}/${OUT_DIR}/api
echo 'build done.'
