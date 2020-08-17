# Papilo
![Go](https://github.com/thealamu/papilo/workflows/Go/badge.svg)

Stream data processing micro-framework; Read, clean, process and store data using pre-defined and custom pipelines.

Papilo packages common and simple data processing functions for reuse, allows definition of custom components/functions and allows permutations of them in powerful ways using the *Pipes and Filters Architecture*.

## Features (WIP)
- **Pipeline stages**: Read - Process - Store
- **Pre-defined pipelines**: Papilo offers pre-defined pipelines and components 
- **Custom pipelines**: Organize components to create a custom pipeline flow
- **Concurrency**: Run multiple pipelines concurrently
- **Extensibility**: Extend by adding custom components
- **Network source**: REST and WebSocket APIs for data ingress
- **Custom source**: Define a custom source for data ingress
- **Multiple formats**: Transform input and output data using transformation functions

### Architecture
![Architecture](./images/architecture.svg)

The Pipes and Filters architectural pattern provides a structure for systems that process a stream of data.
In this architecture, data is born from a **data source**, passes through **pipes** to intermediate stages called **filter components** and ends up in a **data sink**. Filter Components are the processing units of the pipeline, a filter can enrich (add information to) data, refine (remove information from) data and transform data by delivering data in some other representation. Any two components are connected by pipes; Pipes are the carriers of data into adjacent components. Although this can be implemented in any language, Go lends itself well to this architecture through the use of channels as pipes.

### Defaults
Papilo offers default sources, sinks and components:

- Sources:
    - File: Read lines from a file
    - Stdin: Read lines from standard input (default)
    - (WIP): Network: A REST endpoint is exposed on a port
    - (WIP): WebSocket: Full duplex communication, exposed on a port

- Sinks:
    - File: Write sink data to file
    - Stdout: Write sink data to standard output (default)

- Components:
    - Sum: Continuously push the sum of all previous numbers to the sink


## Examples
Read from stdin, write to stdout:
```go
package main

import "github.com/thealamu/papilo/pkg/papilo"

func main() {
	p := papilo.New()
	m := &papilo.Pipeline{} // Default data source is stdin, default data sink is stdout 
	p.Run(m)	
}
```

Add a stream of numbers:
```go
func main() {
	p := papilo.New()
	m := &papilo.Pipeline{
		Components: []papilo.Component{papilo.SumComponent},
	}
	p.Run(m)
}
```

Make every character in stream lowercase using a custom component:
```go
func lowerCmpt(p *papilo.Pipe) {
	for !p.IsClosed { // read for as long as the pipe is open
		// p.Next returns the next data in the pipe
		d, _ := p.Next()
		byteData, ok := d.([]byte)
		if !ok {
			// we did not receive a []byte, we can be resilient and move on
			continue
		}
		// Write to next pipe
		p.Write(bytes.ToLower(byteData))
	}
}

func main() {
	p := papilo.New()
	m := &papilo.Pipeline{
		Components: []papilo.Component{lowerCmpt},
	}
	p.Run(m)
}
```