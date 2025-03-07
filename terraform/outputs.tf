output "ecs_cluster_name" {
  description = "Name of the ECS cluster"
  value       = aws_ecs_cluster.blockchain_cluster.name
}

output "ecs_service_name" {
  description = "Name of the ECS service"
  value       = aws_ecs_service.blockchain_service.name
}

output "vpc_id" {
  description = "ID of the VPC"
  value       = module.vpc.vpc_id
}