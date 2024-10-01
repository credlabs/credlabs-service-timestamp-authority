package main

import (
	"context"
	//"time"
	"fmt"
	//"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event *CredlabsTimestampRequest) (*CredlabsTimestampResponse, error) {
	err := validateTimestampRequest(event)
	if err != nil { return nil, err }

	response := CredlabsTimestampResponse{
		Status: fmt.Sprintf("Hello %s!", event.PublicKey),
	}

	return &response, nil
}

func init() {
	
	fmt.Println("init: Start")

	go startNtpMonitor()
	defer stopNtpMonitor()

	fmt.Println("init: End")
}

func main() {

	fmt.Println("main: Start")
/*
	// Request
	var jsonData = "{\r\n  \"publicKeyType\": \"pem\",\r\n  \"publicKey\": \"-----BEGIN CERTIFICATE-----\\r\\nMIIGSDCCBDCgAwIBAgICEAAwDQYJKoZIhvcNAQELBQAwYTELMAkGA1UEBhMCQVUx\\r\\nDDAKBgNVBAgMA1ZJQzEXMBUGA1UECgwOQ3JlZGxhYnMsIEluYy4xKzApBgNVBAMM\\r\\nIkNyZWRsYWJzIEludGVybWVkaWFyeSBUaW1lc3RhbXAgQ0EwHhcNMjQwODA0MTMy\\r\\nNTE0WhcNMjUwODE0MTMyNTE0WjBUMQswCQYDVQQGEwJBVTEMMAoGA1UECAwDVklD\\r\\nMRcwFQYDVQQKDA5DcmVkbGFicywgSW5jLjEeMBwGA1UEAwwVdGltZXN0YW1wLmNy\\r\\nZWRsYWJzLmlvMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtMdNzy70\\r\\n+JOEoQYE0K4fwffDmMvZCIhE6ipGU+qYlmbOMTqflVFAN2jhC+1AUtipS3HK2nny\\r\\n54eSrkUSjLEy7kO4BhKgoL4mLBWUxytQ07Qntb9e1q8D\\/0MInBLsW\\/YmG3IbVYli\\r\\n2Ap6atKJWzkfrinyLeEpyYcHeruu24Iuxy3JXX4l7ex4NGyv3iYsOeGw01LdRr3E\\r\\nT1W\\/YBZv6d999reA1gsZzCUpedOmKGxEVxXtICqTLxuaMQ0xml12wpgh0lRbfYPM\\r\\nsy0LxF0deVdghrh6D0Jh7lNTm+c5GPohjzT55HfJpUB\\/MQF18ema5XAYyuPr0Yfq\\r\\nnAl6rrx8djwImEEWTOmTI6\\/ephhJ4wMBMb0s1Gke1clLADhKm4SdyfPfs5BcVieq\\r\\n4z1q\\/QbnhzNz+uHia\\/1JCoK+k0Bnj3JIbQetZOIncHAgxOg3Ox+GsHFJ5RHmEJO9\\r\\nGe+JIVVXMjW6UN4cV7a8YPLx3tt4XuNE9qr8ox9I2KT3cVic3jaY77KvkM5aKRqR\\r\\nCy03FD7QHUM690br5c\\/a8jgjnj3oH6BHC2w+5bS+eC8hFiHgBO61Fo1FLTrrESnv\\r\\n8BvGSlIh2GPLwLMbNNPuUXLge39uFDKYGbyCPl2hQ\\/dWfJXKs8pzHI0R9Q9FAe3P\\r\\nUB17Kto2RUY9FkEK9CKXW4RTvvQ0yyF2qcECAwEAAaOCARUwggERMAkGA1UdEwQC\\r\\nMAAwEQYJYIZIAYb4QgEBBAQDAgZAMDMGCWCGSAGG+EIBDQQmFiRPcGVuU1NMIEdl\\r\\nbmVyYXRlZCBTZXJ2ZXIgQ2VydGlmaWNhdGUwHQYDVR0OBBYEFMmBUBjhzvwc82IX\\r\\n1+Q+PyTWgAVkMHgGA1UdIwRxMG+AFP74RKUQ3SdsSKbvvNeYxgk\\/kxQmoVOkUTBP\\r\\nMQswCQYDVQQGEwJBVTEMMAoGA1UECAwDVklDMRcwFQYDVQQKDA5DcmVkbGFicywg\\r\\nSW5jLjEZMBcGA1UEAwwQQ3JlZGxhYnMgUm9vdCBDQYICEAAwCwYDVR0PBAQDAgbA\\r\\nMBYGA1UdJQEB\\/wQMMAoGCCsGAQUFBwMIMA0GCSqGSIb3DQEBCwUAA4ICAQCEF5sn\\r\\n+MiEIoX1RgA0XOIPBZnQwG+VsK8oNyDymosX5K\\/ZXRfnKz0AHEwYsx1fvteSIx\\/3\\r\\ncSWBA7jU4Pg1GF\\/A8h7T9951qtJEzmrr713jpTZi\\/8Xyra4qrLEp20d5xaB4lncv\\r\\nPIe2qjLXOwQ02fiCrliQLgzL4Uh0hdPKK165fTOI3bHqV0QLTHFdigacWXcinNHS\\r\\nIzIjdqiRY4ni6yFO6agbIQGv8IQ9rcyNPiceZuNFJQ7iOQKW7kRl0ZV19ApeijNg\\r\\nUJ6C4bfUmPWO+CpORhQ0xZLmFj90nTodOZpWyr3qzYpIi07GlTn+4kb5fVpEqtkh\\r\\nUtJ9tUyyGvQdfZsAMDLwl3P+BwXlcesn4WyclVX91zXCcDglTcp4GkX3FWbRy73i\\r\\nhOq1HuoF3TOea+O+3P\\/UnnwwsCLJWaqrduL28dPwLPV0QK16rWd3a5qIgts6dlbo\\r\\nr6B2SJcNybQ13baNZNTE1mM2Z5xRWBlkcX8ExReIWPBDDGZ5LH6+\\/0J7H8yXcE6s\\r\\n0PtBRMp4dsjzfoyZijSmvcx3K1u24vqX7akSgLHUQvfyyxs19lN9bTv4doMujcB6\\r\\nTY1Pejp2ytqS4JZFuwWLKj2CHNIv4PfJtQ+2Pbm020A7EUk6nGeOzlbSi0wju0tm\\r\\nNTWepMxXdjapqwJMwkChj\\/3lpocHmLBYO+rjEg==\\r\\n-----END CERTIFICATE-----\",\r\n  \"privateKeyType\": \"kms\",\r\n  \"privateKey\": \"alias/timestamp-key-beta-3\",\r\n  \"certificateChainType\": \"pem\",\r\n  \"certificateChain\": \"-----BEGIN CERTIFICATE-----\\r\\nMIIFkjCCA3qgAwIBAgICEAAwDQYJKoZIhvcNAQELBQAwTzELMAkGA1UEBhMCQVUx\\r\\nDDAKBgNVBAgMA1ZJQzEXMBUGA1UECgwOQ3JlZGxhYnMsIEluYy4xGTAXBgNVBAMM\\r\\nEENyZWRsYWJzIFJvb3QgQ0EwHhcNMjQwODA0MTMyNTEyWhcNMzQwODAyMTMyNTEy\\r\\nWjBhMQswCQYDVQQGEwJBVTEMMAoGA1UECAwDVklDMRcwFQYDVQQKDA5DcmVkbGFi\\r\\ncywgSW5jLjErMCkGA1UEAwwiQ3JlZGxhYnMgSW50ZXJtZWRpYXJ5IFRpbWVzdGFt\\r\\ncCBDQTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAKY9Tz3nbPhhm4Ur\\r\\nCu12dQ6H\\/M\\/mBSfQqfXssAVITK9KUF+fj6AedKfHMeeNbM6ZvZoVv0YB1RLFkydq\\r\\n0TnTbSH08LiRXquZ5pUgrdqutPD5zKFPFSmicZWcoCNyyxpewGd6kBAg3fq804A+\\r\\nf8E\\/C8Wdlza34TuIclwUlmXXRL5fkpLXyKTlVnmS+JLasMlx3N8jYTWQUozwo2uh\\r\\nt4k10nne1QClHGkq6N\\/Iypt8nFO3zX0w73552v2eRVVf+Ohkskf6qCcKIFxD0N3\\/\\r\\nHuKVL9j3uUKGm5EXtbSsF7EsFqGyVJf6bue+9s\\/tG+l5wtK1t6pTHdoTUxDcueDC\\r\\nCZuTxBWu3VY\\/puCyvGSDbo7Ut7ST81CfWsSTWf\\/7RpTEueXBmxdSocGzneIdXG+t\\r\\ngvkfHwH+uMDCjF\\/tKPbnvOTGV0AxB3k9cCT8JDpYUJ+KSMfIm7roYRlr2rmYJPlN\\r\\naQTXi1UfzR\\/zn0PapW0FKgFQg4\\/hRuIolK0bEQpbHxvsMt8iij1LwnllqoAbPyIJ\\r\\nFAaH49NvOPmqD6ChQtuyiIqKk2eOwM8LnckPoNvY94kPuzngUQRDb5QMX0biSflo\\r\\njmofDxTGeWNjWxm9kZ7Dcc8XZKHDCT3uio\\/5g6HUG\\/hP5Rdb5aLGMwu\\/h29jTXaJ\\r\\nf+4Jew8IZLhVSP+dCjfS+MqkVkaXAgMBAAGjZjBkMB0GA1UdDgQWBBT++ESlEN0n\\r\\nbEim77zXmMYJP5MUJjAfBgNVHSMEGDAWgBRGbhdf6QrwObBu6r3RfB4zH0GQUzAS\\r\\nBgNVHRMBAf8ECDAGAQH\\/AgEAMA4GA1UdDwEB\\/wQEAwIBhjANBgkqhkiG9w0BAQsF\\r\\nAAOCAgEAC4v+UwLx+amXE\\/\\/X79JQa9SFWq\\/Mx92RwHoyilfmN3SxN7elfR6o8Drb\\r\\nK\\/zfyA3I5qD64TEtzylrejF30IfVq18cheSZAViqBVjBPtWTfzolRmmpvDLjs5LJ\\r\\n0hjreSml2zRruDSRDokB9QgZXDHePnkvSnMmtBDEPAEieniJ1sTpRQfMCxv58dwS\\r\\nXAOaGyXPZFI4Lk692tpgQBoURZ7xry649Xqm8\\/rxevg+6JlAMqfMRFT9ZdOka9ff\\r\\n5TbMFKjAnRQUnbdiS0cD4nnWU0Q5QgKUmpgvWE3Gfb6jHDiTL3PRDmbAmfHYAgz\\/\\r\\nKS\\/ChWR1gfTIKvt1OWE5pTAHk3a2HLvPFUgSokjx\\/qeQ+vHVSYkvdqHvL46kER5K\\r\\npZF\\/8z3pXJ7wtqIlS22PKode5Pger9GIKAoFhg9xAwYI3ukOdFEm7CBPzutDAbVw\\r\\nSa6pkcbFVwpUEp5NaIpIGFLK2TwSQ\\/KiWY6+Xno+AYWp+Rvm+MsDdZXZmnBnD+Bd\\r\\nIVvZ9MsuCF\\/QAxppkhE8ZklZssb5kDP7R6arRw1We3j8p0bveNW5JR8qSBekUFhd\\r\\nJ59TG0OLgdZ2uYFiSTQ\\/KM+Rh7NNzAu9QfuqS4zPZ04BsCUUbleGT4e0RaAo7V4p\\r\\nQpNydwgvkEJY4zym\\/6tpU9fjfla3m4O27+HeSHF8QeA2PIHtugo=\\r\\n-----END CERTIFICATE-----\\r\\n-----BEGIN CERTIFICATE-----\\r\\nMIIFjzCCA3egAwIBAgIUG3aBQomRqxzOMbr8MLei0924CeQwDQYJKoZIhvcNAQEL\\r\\nBQAwTzELMAkGA1UEBhMCQVUxDDAKBgNVBAgMA1ZJQzEXMBUGA1UECgwOQ3JlZGxh\\r\\nYnMsIEluYy4xGTAXBgNVBAMMEENyZWRsYWJzIFJvb3QgQ0EwHhcNMjQwODA0MTMy\\r\\nNTEwWhcNNDQwNzMwMTMyNTEwWjBPMQswCQYDVQQGEwJBVTEMMAoGA1UECAwDVklD\\r\\nMRcwFQYDVQQKDA5DcmVkbGFicywgSW5jLjEZMBcGA1UEAwwQQ3JlZGxhYnMgUm9v\\r\\ndCBDQTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAJfPNDSKjJNuM+7d\\r\\n565Jox5v9hfz15qmpfjLvUVAdwo46hufshKhnW\\/xzsz2adhNxi8BHXN\\/MdGuOEOp\\r\\nug6Ii\\/U9LHj1WnBGarUO3D8F8\\/auUYGYffNhI067Fwy9XfpWbsHrsMQ589T29Sr+\\r\\nRB151ZaOBQXGiZz6pvCK2jxwD2kN3VO2oBOgC5NINUKE\\/vlV\\/ZxCRPKr2M51Pp0K\\r\\n167ACZ2pkgGrtrxDj66ZWd1MWGWOEnNXNdMn3ooQBdWW9\\/0xmQHCqm+bxHXuzApK\\r\\n1SfIDfmutk9QAt5sDC8AGGrsSuYGGW5NocP+8xHnKOCbctgKqbIFsPQtfh67\\/VE\\/\\r\\nzkpJPwB60BC0bc63c5lwx07qp0A3Rwe1U3rk+GS4KTP6+Pr8b43raVCmZ4Kutsmt\\r\\nmGAqcYAQseg9oPh\\/AL00JRDSTIyCu+xAVbBmQHvj\\/Fy3JFWQIU7vKNYqzQ5hHli0\\r\\nuLba9qO3Wphsezvz18DXOgUN2JJxyV7pIuNq5SpLwxt8de96PU0WH6sY2Ceuhn9b\\r\\n\\/mAme8XpkHWEx9AVXEYt0z+zJ0\\/ozg68MeKgyGCVniGp496NYSx2iiAYwhapSUEP\\r\\niwJhVuuLhY5ro7M0\\/xaHTeQ4LzRjYK6MDi6vrNZz+Kwzbz7FyPZCceDVR6F7kuwJ\\r\\n1JDi4noI\\/qUh3EXraE+mLbQfLHLPAgMBAAGjYzBhMB0GA1UdDgQWBBRGbhdf6Qrw\\r\\nObBu6r3RfB4zH0GQUzAfBgNVHSMEGDAWgBRGbhdf6QrwObBu6r3RfB4zH0GQUzAP\\r\\nBgNVHRMBAf8EBTADAQH\\/MA4GA1UdDwEB\\/wQEAwIBhjANBgkqhkiG9w0BAQsFAAOC\\r\\nAgEAir44eC6Nvma9NOfT1DgUnrXBuhsI44Snon2CEBnLMDXshW3cZJuLgFSpmdXv\\r\\n6fLeNtDH4qiZs8UFmJdv1kTS6o2YYQ6RjNgjSLQgovA7CWs+XqbrCCFmcDqv+t9E\\r\\nJHy\\/YvuJtstGhVgb5yEHt\\/2Tu3MeO7+AABd8E\\/zeJMhkixEvpVy+u\\/reCLWjkKR2\\r\\nUvOI7TAqEQ3JY8Km03kzDn8+spnNWhsbe9gi2h5MEX4Yuc4UlI7vueVlC3Ba0wsg\\r\\neEBDopDKp9V81+J19OrkNUbNJDBJmtX8gF1Ye5aQXqej0vu4TgmcZJ2VGzO2XFty\\r\\ntNM73Q8J9DdrWaAaxxOT9W0usFmVT80fwFkZAAtfUpalMf+lA633yL0tf5CswlB3\\r\\nIVQ3Wi6IbgHLMxtDZCZV5Q\\/OP\\/vE9q0RaEzxSlthdOnA38F86tjCDeClysRabZrf\\r\\nauXUde7lkr2lLLYeGSYdRxEdYqXBId9zO7gqJk+p3Nw61QePUUjjC9reiuaZf+Ra\\r\\nJNsuUx79h2hRR7lOW6HqT2K6fkJLW3N30l1qUORYAr6O9egy9gfzPEOf7URt13zz\\r\\n\\/w2VzxDyb3hYuXYWbWR4I5eUJab9kHPvWxoIXGrXE7aHsMElvaPx5sCYFl9so6Zk\\r\\nGK6rDCNqVHNJTwpopN8+W0Ko\\/CfO\\/+wN6yQQUttXVht06\\/g=\\r\\n-----END CERTIFICATE-----\",\r\n  \"request\": \"MGMCAQEwUTANBglghkgBZQMEAgMFAARAkA5DzFXWuM3PaMilEDMNJFsvaIs4nq9IiNC7UGw/fgzTetW6W1sgYNp8ldssH1AegMcFeJ+Q39icmBJ7IHduRQIICHShvc6DbwsBAf8=\"\r\n}";
	
	var event CredlabsTimestampRequest
	json.Unmarshal([]byte(jsonData), &event)

	response, err := HandleRequest(nil, &event)
	if err != nil { panic(err) }

	out, err := json.Marshal(response)
	if err != nil { panic(err) }

	fmt.Println(string(out))
*/
	// --------------------------------

	lambda.Start(HandleRequest)
/*
	// Calling Sleep method 
	time.Sleep(8 * time.Second) 
*/
	fmt.Println("main: End")
}