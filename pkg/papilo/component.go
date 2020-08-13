package papilo

type Component func(in chan []byte, out chan []byte)
