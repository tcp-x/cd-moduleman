# get current repository latest version
echo "current repository latest version:\n"
git ls-remote --tags https://github.com/tcp-x/cd-moduleman.git
# cd-cli plugin compile 
cd ./moduleman/
go build -buildmode=plugin -o ../plugins/Moduleman.so moduleman.go
cd ../cd-obj/
go build -buildmode=plugin -o ../plugins/CdObj.so cd-obj.go
cd ../repo/
go build -buildmode=plugin -o ../plugins/Repo.so repo.go
cd ..

# go build -buildmode=plugin -o Moduleman.so
# set latest version
Version="v0.0.5"
go mod tidy
git submodule update --remote
git add go.mod moduleman/moduleman.go cd-obj/cd-obj.go repo/repo.go repo/repo-tag.go plugin/Moduleman.so plugin/Repo.so plugin/CdObj.so
git add -A
git commit -a -m "set version $Version"
git tag $Version
git push origin $Version



