param (
	[int]$End = 4000000
	[int]$Divisor = 2
)

$Fibonacci = [System.Collections.Generic.List[int64]]::new()

for ( $i = $Start; $i -le $End; $i +- $i ) {
}
