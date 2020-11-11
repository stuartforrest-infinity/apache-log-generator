### About
This project contains a spec and code that will stream dummy Apache logs into a single or multiple Kinesis streams. You must have (locally) AWS credentials that will permit this.

### Usage
  - Build the container - `docker build --tag apache-log-mutiple .`
  - Run the container
    - Replace the `AWS_PROFILE` value environment arg in the command below to the local AWS credentials profile that has permissions to put records into the streams you are targeting
    - With a single stream using the `--stream` flag - `docker run -it -v ~/.aws:/root/.aws -e AWS_PROFILE=your-profile-name -e AWS_REGION=eu-west-1 apache-log-mutiple:latest -stream=stream-name`
    - With single or multiple streams using the `--streams` flag - `docker run -it -v ~/.aws:/root/.aws -e AWS_PROFILE=your-profile-name -e AWS_REGION=eu-west-1 apache-log-mutiple:latest -streams=stream-name-1,stream-name-2`

