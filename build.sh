# Get version from git

# Get the latest Git tag
	ApiVers=$(git describe --tags --abbrev=0)

# Get the current Git commit and branch
	Commit=$(git rev-parse --short HEAD)
	Branch=$(git rev-parse --abbrev-ref HEAD)

# Get the build date
	BuildDate=$(date +'%Y-%m-%d %H:%M:%S %Z')

# Get Go version and OS/Arch
	GoVers=$(go version | { read _ _ v _; echo $v; })
	OsArch=$(go version | { read _ _ _ v; echo $v; })

# Set the LDFlags variable
	LDFlags="-s -w \
	-X 'main/utils.ApiVersion=$ApiVers' \
	-X 'main/utils.Commit=${Branch}-${Commit}' \
	-X 'main/utils.BuildDate=$BuildDate' \
	-X 'main/utils.GoVersion=$GoVers' \
	-X 'main/utils.OsArch=$OsArch'"

# Build your Go program with the specified LDFlags
	CGO_ENABLED=0 go build -v -ldflags="$LDFlags" -o grpc-gateway-server