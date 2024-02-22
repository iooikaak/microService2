package enum

type JaegerConfig float64

const (
	JaegerSampleConfigParam JaegerConfig = 1
)

func (j JaegerConfig) Float64() float64 {
	return float64(j)
}

type Jaeger string

const (
	JaegerStartSpan Jaeger = "microService2 start tracing"
)

func (j Jaeger) String() string {
	return string(j)
}
