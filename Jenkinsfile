
stage ('build deploy') { 
    node{
       sh """
       rm -rf ./src/*
       """
     
    
       dir ('src/github.com/chj050322/'){
          git credentialsId: '51c867ae-e565-41d7-8899-f8e278089a56', url: 'https://github.com/chj050322/test.git'
       }
	   
	    dir ('src/github.com/jingyan/adonis'){
          git credentialsId: '51c867ae-e565-41d7-8899-f8e278089a56', url: 'https://github.com/jingyan/adonis.git'
       }
       
           
      sh """

      source /etc/profile
      export GOPATH=$WORKSPACE
      mkdir -p $WORKSPACE/bin
      export GOBIN=$WORKSPACE/bin

      cd $WORKSPACE/src/github.com/jingyan/adonis
      
      rm -rf vendor/* 2>/dev/null
      cp -r /var/lib/jenkins/go-depen-pkg/vendor/src/*  vendor/ 
      
      cd $WORKSPACE/src/github.com/jingyan/adonis
      go build  -o edz-adonis .
      zip -j edz-${JOB_NAME}.zip edz-adonis
	  
	  mkdir -p  /home/jenkins/run/config
	  rm -rf /home/jenkins/run/config
	  cp edz-adonis  /home/jenkins/run
	  cp -r config   /home/jenkins/run
		
	  cd  /home/jenkins/run
	  nohup  /home/jenkins/run/edz-adonis  >>/home/jenkins/log/edz-adonis.log 2>&1 &
		
     
      """
       
    

  
   }
}