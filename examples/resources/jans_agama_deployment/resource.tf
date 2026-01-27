resource "jans_agama_deployment" "example" {
  name = "example-agama-project"
  
  deployment_file = "path/to/agama-project.gama"
  
  autoconfigure = true
}
