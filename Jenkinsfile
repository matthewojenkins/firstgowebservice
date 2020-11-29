 pipeline {
    agent any

    stages {
        stage('Latest code checkout') {
            steps {
                echo 'Pulling latest service code'
                checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/matthewojenkins/firstgowebservice.git']]])
            }
        }
        stage('Building service and docker image') {
            steps {
                echo 'Building service and docker image'
                bat 'build.bat'
            }
        }
    }

 }