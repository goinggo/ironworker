export HOME_FOLDER="$HOME/Container"
export CODE_FOLDER="$HOME_FOLDER/code"
export PROGRAM_FOLDER="$CODE_FOLDER/src/github.com/goingo/ironworker"

mkdir $HOME_FOLDER
cd $HOME_FOLDER
curl https://go.googlecode.com/files/go1.1.2.linux-amd64.tar.gz -o p.tar.bz2 && tar xf p.tar.bz2 && rm p.tar.bz2
export GOARCH="amd64"
export GOBIN=""
export GOCHAR="6"
export GOEXE=""
export GOHOSTARCH="amd64"
export GOHOSTOS="linux"
export GOOS="linux"
export GOPATH="$CODE_FOLDER"
export GORACE=""
export GOROOT="$HOME_FOLDER/go"
export GOTOOLDIR="$HOME_FOLDER/go/pkg/tool/linux_amd64"
export CC="gcc"
export GOGCCFLAGS="-g -O2 -fPIC -m64 -pthread"
export CGO_ENABLED="1"
export PATH=$GOROOT/bin:$PATH

go get -x github.com/goinggo/ironworker