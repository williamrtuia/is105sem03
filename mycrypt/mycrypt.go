package mycrypt

var ALF_SEM03 = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
        kryptertMelding := make([]rune, len(melding))
        for i := 0; i < len(melding); i++ {
                indeks := sokIAlfabetet(melding[i], alphabet)
                if indeks == -1 {
                        kryptertMelding[i] = melding[i]
                        continue
                }
                if indeks+chiffer >= len(alphabet) {
                        kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]
                } else {
                        kryptertMelding[i] = alphabet[indeks+chiffer]
                }
        }
        return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
        for i, s := range alfabet {
                if s == symbol {
                        return i
                }
        }
        return -1
}

