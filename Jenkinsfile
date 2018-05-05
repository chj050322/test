
stage ('build deploy') { 
    node{
       sh """
       rm -rf ./src/*
       """
     
    
       dir ('src/github.com/jingyan/adonis'){
          git credentialsId: '51c867ae-e565-41d7-8899-f8e278089a56', url: 'https://github.com/chj050322/test.git'
       }
       
       withCredentials([usernameColonPassword(credentialsId: 'e11cefdb-5deb-4965-8e6b-68e6927f1658', variable: 'hk-ssh')]){       
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

      if [[ ${rsync_salt} == "true" ]]
      then
        rsync -ravz -e "ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${hk-ssh}"  --progress edz-${JOB_NAME}.zip  ${hk-ssh}:~/test/run/adonis/
        ssh -tt -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${hk-ssh} ${hk-ssh} "sudo salt '*adons*' state.highstate"
      fi
      """
       
    }

    withCredentials([usernameColonPassword(credentialsId: '51c867ae-e565-41d7-8899-f8e278089a56', variable: 'TOKEN')]) {
       sh """
       if [[ ${release} == "true"  &&   -v releaseVersion ]]
       then
       export GITHUB_TOKEN=${TOKEN}
       cd $WORKSPACE/src/github.com/edz/${JOB_NAME}

       
       fi    
       """       
    }
   }
}