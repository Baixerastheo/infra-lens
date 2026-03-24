resource "aws_instance" "api" {
  ami           = "ami-12345"
  instance_type = "t2.micro"
}

resource "aws_security_group" "web" {
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_s3_bucket" "assets" {
  bucket = "my-assets"
}

resource "aws_db_instance" "main" {
  engine                  = "mysql"
  instance_class          = "db.t2.micro"
  backup_retention_period = 0
}