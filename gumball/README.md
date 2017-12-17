##### Description

The gumball machine was deployed to dockercloud and has a design flaw when over a distributed deployment.

Check the docker stack file to see the distributed deployment.

Also check the dependencies : github.com/
  -codegangsta/
  -gorilla/
  -satori/
  -streadway/
  -unrolled/



  #####  Setting GO Path:

  `export GOPATH="$(pwd)"
  source ~/.bash_profile
  echo $GOPATH
  go build gumball
  ./gumball`


  ##### Build Docker Image for goapi code

  `docker build -t gumball:latest .
  docker login
  docker tag gumball:latest nevosial/gumball:v1
  docker push nevosial/gumball:v1`
