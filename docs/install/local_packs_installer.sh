#Set this to your $GOPATH environment variable's value
#In my case, I'm setting it to   /Users/admin/projects/Go/GoWorkspace
GOPATH="/Users/admin/projects/Go/GoWorkspace"
GOIRIS=$GOPATH/src/goiris
cat ./synkku_print.txt
echo ""
echo "****************************************************************"
echo "..********************** FORMATTING **************************.."
echo "..**********************  BUILDING ***************************.."
echo "..********************** INSTALLING **************************.."
echo ""
cd $GOPATH/src/goiris/conf
go build
go install 
go fmt
echo "1) conf successfully installed"

cd $GOPATH/src/goiris/views 
go build
go install 
go fmt
echo "2) views successfully installed"

cd $GOPATH/src/goiris/controllers
go build
go install 
go fmt
echo "3) controllers successfully installed"

cd $GOPATH/src/goiris/routers/
go build
go install 
go fmt
echo "4) routers successfully installed"

cd $GOPATH/src/goiris/
go build
go fmt main.go
echo "5) synkku successfully built (synkku.exe modified)"
echo ""

for ((i=1;i<74;i++))
do
	printf "#"
	sleep .02
done
echo ""
# echo "PACKAGES(conf,views, controllers, routers, goiris) SUCCESSFULLY INSTALLED"
# echo ""
# echo ""

#END

# echo ""
# echo ">>> Installing conf..."
# go fmt $GOIRIS/conf/
# echo $GOIRIS"/conf/ is successfully formatted..."
# go build $GOIRIS/conf/
# echo $GOIRIS"/conf/ is successfully built..."
# go install $GOIRIS/conf/
# echo $GOIRIS"/conf/ is successfully installed..."

# echo ""
# echo ">>> Installing views..."
# go fmt $GOIRIS/views/
# echo $GOIRIS"/views/ is successfully formatted..."
# go build $GOIRIS/views/
# echo $GOIRIS"/views/ is successfully built..."
# go install $GOIRIS/views/
# echo $GOIRIS"/views/ is successfully installed..."

# echo ""
# echo ">>> Installing controllers..."
# go fmt $GOIRIS/controllers/
# echo $GOIRIS"/controllers/ is successfully formatted..."
# go build $GOIRIS/controllers/
# echo $GOIRIS"/controllers/ is successfully built..."
# go install $GOIRIS/controllers/
# echo $GOIRIS"/controllers/ is successfully installed..."

# echo ""
# echo ">>> Installing routers..."
# go fmt $GOIRIS/routers/
# echo $GOIRIS"/routers/ is successfully formatted..."
# go build $GOIRIS/routers/
# echo $GOIRIS"/routers/ is successfully built..."
# go install $GOIRIS/routers/
# echo $GOIRIS"/routers/ is successfully installed..."


# echo ""
# echo ">>> Installing goiris..."
# go fmt $GOIRIS
# echo $GOIRIS"/ is successfully formatted..."
# go build $GOIRIS
# echo "goiris.exe successfully created..."

