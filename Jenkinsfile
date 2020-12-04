node {
    def app

    stage('Clone repository') {

        checkout scm
    }

    stage('Test Code') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("mintojoseph/mweb:test", "--target test .")
    }

    stage('Build image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("mintojoseph/mweb")
    }

    stage('Push image') {
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
        }
    }
    stage('Set Terraform path') {
        script {
         def tfHome = tool name: 'Terraform'
         env.PATH = "${tfHome}:${env.PATH}"
        }
        sh "terraform -version"
    }
 
    stage('Terraform Init') {
      
        sh "cd terraform; terraform init -input=false"
      
    }
    stage('Terraform Plan') {
      
        sh "cd terraform; terraform plan -out=tfplan -input=false"
      
    }
    stage('Terraform Apply') {
      
        sh "cd terraform; terraform apply -input=false tfplan"
      
    }
}
