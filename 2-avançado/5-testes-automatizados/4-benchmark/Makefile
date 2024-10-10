# usaremos o graphviz para gerar os gráficos https://www.graphviz.org/download/
# para ver os gráficos de memória e cpu, execute os comandos abaixo:
# make bench
# make mem
# make cpu

bench:
	go test -bench=. -benchmem -run=^$ . -memprofile memprofile.out -cpuprofile cpuprofile.out

mem:
	go tool pprof -web memprofile.out

cpu:
	go tool pprof -web cpuprofile.out

mem-text:
	go tool pprof -text memprofile.out

cpu-text:
	go tool pprof -text cpuprofile.out