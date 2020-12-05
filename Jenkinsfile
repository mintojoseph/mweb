node {
    def app

    stage('Clone repository') {

        checkout scm
    }

    stage('Test Code') {

        app = docker.build("mintojoseph/mweb:test", "--target test .")
    }

    stage('Build image') {

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
      
        sh "cd deployment/terraform; terraform init -input=false"
      
    }

// Not the right way. Correct way would be to use a remote state. 
// Destroying for workarounding state issues.
    stage('Destroy') {
        job:
          script:
            - false || exit_code=$?
            
        sh "helm uninstall mweb || exit_code=$?"

    }

    stage('Terraform Plan') {
      
        sh "cd deployment/terraform; terraform plan -out=tfplan -input=false"
      
    }
    stage('Terraform Apply') {
      
        sh "cd deployment/terraform; terraform apply -auto-approve -input=false tfplan "
      
    }
}
