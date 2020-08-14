# Papilo
![Go](https://github.com/thealamu/papilo/workflows/Go/badge.svg)

Stream data processing micro-framework; Read, clean, process and store data using pre-defined and custom pipelines.

Papilo packages common and simple data processing functions for reuse, allows definition of custom components/functions and allows permutations of them in powerful ways using the *Pipes and Filters Architecture*.

## Features
- **Pipeline stages**: Read - Clean - Process - Store
- **Extensibility**: Extend by adding custom components
- **Pre-defined pipelines**: Papilo offers pre-defined pipelines in its command line tool
- **Custom pipelines**: Organize components to create a custom pipeline flow
- **Concurrency**: Run multiple pipelines concurrently
- **Network source**: Papilo exposes REST and WebSocket APIs for data ingress
- **Custom source**: Define a custom source for data ingress
- **Multiple formats**: Transform input and output data using transformation components

### Architecture
![Architecture](./images/architecture.svg)

The Pipes and Filters architectural pattern provides a structure for systems that process a stream of data.
In this architecture, data is born from a **data source**, passes through **pipes** to intermediate stages called **filter components** and ends up in a **data sink**. Filter Components are the processing units of the pipeline, a filter can enrich (add information to) data, refine (remove information from) data and transform data by delivering data in some other representation. Any two components are connected by pipes; Pipes are the carriers of data into adjacent components. Although this can be implemented in any language, Go lends itself well to this architecture through the use of channels as pipes.

### Defaults
Papilo offers default sources, sinks and components