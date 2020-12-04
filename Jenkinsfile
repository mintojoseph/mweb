node {
    def app

    stage('Clone repository') {

        checkout scm
    }

    stage('Test Code') {

            agent {
                dockerfile {
                    filename 'Dockerfile'
                    additionalBuildArgs '--target test'
                }
            }
    }

    stage('Build image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("getintodevops/hellonode")
    }

    stage('Push image') {
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
        }
    }
}
