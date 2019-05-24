param (
	[int[]]$Multiples = @(3,5),
	[int]$Start = 1,
	[int]$End = 999
)

$Digits = [System.Collections.Generic.List[int]]::new()

for ( $i = $Start; $i -le $End; $i++ ) {
	foreach ( $ThisMultiple in $Multiples ) {
		if ( $i % $ThisMultiple -eq 0 ) {
			$Digits.Add($i)
			break
		}
	}
}

$Digits | Measure-Object -Sum | Select -ExpandProperty Sum

