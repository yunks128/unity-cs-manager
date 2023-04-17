resource "aws_eip" "ip-test-env" {
  instance = aws_instance.unity-ec2-instance.id
  vpc      = true
}

resource "aws_instance" "unity-ec2-instance" {
  ami = "abc"
  instance_type = "t3.xlarge"
  key_name = var.ami_key_pair_name
  tags {
    Name = var.ami_name
    Deployment = "unity-demo"
  }
  ebs_block_device {
    device_name = "/dev/sda1"
    volume_size = 50
  }
}
