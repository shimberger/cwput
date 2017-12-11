# cwput - Standalone binary to write cloudwatch metrics

I know exists, I needed it anyway.. again.

## Usage

It uses the standard s3cli authentication methods.

	cwput --region=<region> --metric=LoadAverage --namespace="CustomNamespace" --unit=Count --value=1.0 --dimensions="Server=$(hostname -s)"

You can use environment variables to provide secret key & access key:

	AWS_SECRET_KET=abc AWS_ACCESS_KEY=def cwput ...

## Developing

Build the binary:

	go build -o cwput *.go

Run the script:

	./examples/collect-metrics.sh


Or both in one:

	go build -o cwput *.go && ./examples/collect-metrics.sh