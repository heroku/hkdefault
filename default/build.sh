newver=$1

[ -z $newver ] && echo "usage: $0 <VERISON>" && exit 1

eval $(go tool dist env)
go build
cp default $GOOS-$GOARCH-hkdefault-$newver
gzip -9 $GOOS-$GOARCH-hkdefault-$newver
