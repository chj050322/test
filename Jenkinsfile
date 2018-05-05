
stage ('build deploy') { 
    node{
       sh """
       rm -rf ./src/*
       """
     
    
       dir ('src/github.com/jingyan/adonis'){
          git credentialsId: '51c867ae-e565-41d7-8899-f8e278089a56', url: 'https://github.com/chj050322/test.git'
       }
       
           
      sh """

      export WORKSPACE=/var/lib/jenkins/workspace
      export GOPATH=$WORKSPACE
      mkdir -p $WORKSPACE/bin
      export GOBIN=$WORKSPACE/bin

      cd $WORKSPACE/src/github.com/jingyan/
      mkdir -p vendor
      rm -rf vendor/* 2>/dev/null
      tar -C vendor -xzf /var/lib/jenkins/jingyan/gopath/godeps.go-commons.tgz
      mv vendor/src/* vendor/
      cd $WORKSPACE/src/github.com/jingyan/${JOB_NAME}
      go build -ldflags  -i -o edz-$JOB_NAME .
      zip -j edz-${JOB_NAME}.zip edz-${JOB_NAME}

     
      """
       
    

  
   }
}