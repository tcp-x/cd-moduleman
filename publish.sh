# get current repository latest version
echo "current repository latest version:\n"
git ls-remote --tags https://github.com/tcp-x/cd-moduleman.git
# cd-cli plugin compile 
go build -buildmode=plugin -o Moduleman.so
# set latest version
Version="v0.0.2"
go mod tidy
git submodule update --remote
git add go.mod modulenam.go cd-obj.go repo.go repo-tag.go Moduleman.so
git add -A
git commit -a -m "set version $Version"
git tag $Version
git push origin $Version



