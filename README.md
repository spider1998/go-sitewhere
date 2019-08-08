# go-sitewhere


![](https://github.com/spider1998/go-sitewhere/blob/master/sites.svg)


1、安裝Docker CE
--

2、安裝Docker Compose
--
    下載 sudo curl -L “https://github.com/docker/compose/releases/download/1.22.0/docker-compose-(uname -s)-(uname−s)−(uname -m)” -o /usr/local/bin/docker-compose 

    修改許可權 sudo chmod +x /usr/local/bin/docker-compose 

    測試安裝成功否 docker-compose --version docker-compose version 1.22.0, build f46880fe

3、啟用Swarm mode
--
    docker swarm init

4、構建
--
    GRADLE_HOME=/usr/local/gradle-4.10.2 export PATH={GRADLE_HOME}/bin:GRADLEH​OME/bin:{PATH}

    過載/etc/profile這個檔案 source /etc/profile 

    測試 gradle -version 生成gradlew cd /usr/local/gradle-4.10.2 gradle wrapper ln -s gradlew ./bin/gradlew

    下載sitewhere git clone https://github.com/sitewhere/sitewhere.git cd sitewhere git checkout --force sitewhere-2.0.rc2 
    
    執行Gradle構建指令碼前先編輯gradle.properties，
    
    根據自己設定修改 
    
        dockerProtocol=tcp 
        
        dockerHostname=localhost (在本機所以不用ip或域名) 
        
        dockerPort=2375 
        
        dockerRepository=docker.io 
        
        registryUrl=https://index.docker.io/v1/ 
        
        registryUsername=使用者名稱 
        
        registryPassword=密碼 
        
        registryEmail=郵箱

    開始構建映象 gradlew clean dockerImage
     
    如果遇到Could not find tools.jar. Please check that /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.171-8.b10.el7_5.x86_64/jre contains a valid JDK installation cd / find -name tools.jar 
    
    找下看有沒 好吧，可能是閹割過的卸了重新安裝 yum -y remove java-1.8.0-openjdk* yum -y install java-1.8.0-openjdk*

    沒問題的話完工 docker images 應該可以看到構建映象清單 

5、部署
--
    下載 git clone https://github.com/sitewhere/sitewhere-recipes.git 先構建啟動基礎服務映象 cd sitewhere-recipes/docker-compose/infrastructure_defaul docker-compose up 再啟用預設微服務 cd sitewhere-recipes/docker-compose/default docker-compose up

    全部做完沒報錯的話，下載sitewhere-admin-ui 輸入使用者名稱admin密碼password ，IP地址然後登入


*参考文档： https://sitewhere.io/docs/2.0.0/*

*rest API参考文档： https://sitewhere.io/docs/2.0.0/api2/*

*Github地址: https://github.com/sitewhere/sitewhere*
   
   
