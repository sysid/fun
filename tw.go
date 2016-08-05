// multiline string MUST not start wit empty line
package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/gosuri/uilive"
	"github.com/thomas/tw"
	. "github.com/thomas/tw/basic"
)

func init() {
	Log("")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//const letterBytes = ".X"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

var (
	tw1 = `/yyyyyyyyyyyyyyo--:::/+:..'-/++oooooo++++////:----.....----:/:::-------:osyyhyyhhhhhhhhhhhhhhhhhhhh+
ommmmmmmmmmmmmmh::://o/--.-oyhddddddddddddhhyso/:::---:/+oyhhhhhyysooo//shmmmmmmmmmmmmmmmmmmmmmmmmNo
ommmmmmmmmmmmmmd/:::/+---./yhhddhhyhhyysosyyyhys+/:---:/+oossysyyyyyyyysshmmmmmmmmmmmmmmmmmmmmmmmmNo
ommmmmmmmmmmhoyhs+/://-:-+yyyyyhddmmNys/oo+oyhhhs/:--:////+oydhddhhysssyyhmmmmmmmmmmmmmmmmmmmmmmmmNo
ommmmmmmmmmmssyyyhs//::::sysssyhyssyys/:/+oooyhhs+/:::/::///ohdhs/yhyo+oshshmmmmmmmmmmmmmmmmmmmmmmNo
ommmmmmmmmmmhyhsyhds--::+ssoooossooo++////+osyyyso/::-::::://///++++++//o//ommmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmhsoyshdo.-::+ooo+////////////+ossyyyso/:--:/:::::::::::::::://+smmmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmms+hdhs:-:--/+o++///::::::://+osyyyyso/:--:/::------------::///hmmmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmmyoyhdo-:/::/oo+++//::::::://+osyyyyso/:--::::------------:://ommmmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmmdoyhyy:/o/:+ssoo++//:::::///+oyhhyyo+/:--:::--...-.------::/:ommmmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmmmssysh+/oo/ossso++//////////+oyhysso+:---:::--....------:://:ymmmmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmmmmo/yhh++++oyysso++//////://+shhhyyo/:-..-:::....------::://ommmmmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmmmmdoyhds+++osysso+++////:::/+oyhdmmho/:::///-----------:://:dmmmmmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmmmmmmhyhd+++osyysso+++///:///+osyhddhyso+///::----------:///ymmmmmmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmmmmmmmmddo++osyyyssoo+///++oooossssso++o+/////::::::--:://sdmmmmmmmmmmmmmmmmmmmmmmmmmmmo
ommmmmmmmmmmmmmmmmmmmy+++syyyysso+oo++osssssssoo///o+///////::://:://+dmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
oNmmmmmmmmmmmmmmmmmmmy+++syyhhysosso+osyysssoo+////+/:////++/////:://smmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
oNmmmmmmmmmmmmmmmmmmmh+oosyyyhyysso+ooydhyysoo++/:::/:////+++++/////+dmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
oNmNNmNmmmmmmmmmmmmmmh+oosyyyyhyss++/+osyhhyyyssssoooo++++/::://:///ommmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
oNNNNNNNNNmmmmmmmmdmNo+ssyyhyyhyyso++++osssoo+/////::::::/:::://////hmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
oNNNNNNNNNNNNmmmmhmMs++hdyyhyyhhyysooooossyysso+///::::///:::://++/smmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
oNNNNNNNNNNNmNmhymNm:+odmmhhhhyhhhysssoosssysssoossoo+++///:///++/smmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
oNNNNNNNNNNNmhsydNNy:+hdmmmdhhhhhhhyyyssoooooooo++++++////////+++ymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
oNNNNNNNNNdhhhhhmNNh/ohddmmmmdddhhhhyysssoo++////::::::://///++ohNNNmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
oNNNNmmdhyhddddmNNNh++yhdddmmmmddddhhyyyssoo+/::::::::://///+oomNNNNmmmNNNNNmmmmmmmmmmmmmmmmmmmmmmmo
omddddmddmNNmmmNNNNds++yhhhdddmmmddddhhyysso+//:////:::/++++o//mmNNNNNmmmmmmNNNNNNNNNmmmmmmmmmmmmmmo
+ddmNNNNNNNNNNNNNNNmyso+shhhhhdddddddddhhysoo++////////+++oos.+mmNNNNNNmmmmmmmmmNNNNNNNNNNNNNNmmmmmo
oNNNNNNNNNNNNNNNNNNNhsoo++syyhhhhhdddddddhysoo++++++++++oo+s/.ommmNNNNNNmmmmmmmmmmmNNNNNNNNNNNNNNNNo
oNNNNNNNNNNNNNNNNNNNmsoo++/+oyyyyhhhhhhhhhhyyyyssssssooo++s/.-ommmmNNNNNNmmmmmmmmmmmmNNNmmNNNNNNNNNo
oNNNNNNNNNNNNNNNNNNNNdo+++//:/+ssyyyyhyyyyysssoooooo+++/++.'.:ommmmNNNNNNNmmmmmmmmmmmmmmmNNmmmmNNNNo
+hhhhhhhhhhhhhhhhhhhhhs//:::----/++oooooooo+++///////://.'''.-/yyyyhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh+
`

	tw2 = `y//////////////oddhhhyshmmNdyssoooooossssyyyyhddddmmmmmddddhyhhhdddddddho+//://::::::::::::::::::::s
o..............:hhhyyoyddmdo/:------------::/+oyhhhdddhyso/::::://+oooyy+:........................'o
o..............-yhhhysdddmy/::--::::://oo+///:/+syhdddhysoo++/+////////++:........................'o
o...........:o/:+syhyydhds/////:--..'/+yooso/:::+yhddhyyyyso/-:--::/+++//:........................'o
o...........++///:+yyhhhh+/+++/:/++//+yhyyooo/::+syhhhyhhyyyo:-:+y/:/oso+:+:......................'o
o...........:/:+/:-+ddhhs++oooo++ooossyyyyso+///+oyhhdhyhhhyyyyyssssssyyoyyo.......................o
o...........:+o/+:-omdhhsooosyyyyyyyyyyyyso++///+oyhddhyhhhhhhhhhhhhhhhhyys+.......................o
o............+s:-:+hdhddysossyyyhhhhhhhyyso+////+oyhddhhhhddddddddddddhhyyy:.......................o
o............/o/:-odhyhhyoosssyyhhhhhhhyyso+////+oyhddhhhhddddddddddddhhyyo........................o
o............-o/:/+hyoyhs++oossyyhhhhhyyyso/:://osyhddhhhddmmmdmmdddddhhyho........................o
o.............++/+:syooyo+++ossyyyyyyhyyyso/:/++oshdddhhhddmmmmddddddhhyyh/........................o
o.............-oy/::sssso//+oossyyyyyyhyys+::://oyhdmmdhhhmmmmddddddhhhyyo.........................o
o..............-o/:-+ssso+/++ossyyyyyhhhyso/:-..:oyhhhyyydddmddmddddhhyyh-.........................o
o................:/:-ssso+//++osssyyyhyyyso+/:--:/+osyyyhhddddddddddhyyy/..........................o
o..................--osso+///++ossyyyssoooo+++++ossosyyyyyhhhhhhddhhyy+-...........................o
o..................../sss+////++osoosso+++++++ooyyyosyyyyyyyhhhyyhhyys-............................o
o'.................../sss+//::/+o++oso+//+++oosyyyysyhhyyyyyyyyyyhhyy+.............................o
o'...................:soo+///://++osoo/-://+oossyhhhhhyyyysssssyyhyys-.............................o
o'''.................:soo/////:/++ssyso+/::///++++oooossssyhhhyyhyyyo..............................o
o'''''''''........-.'os++//://://+osssso+++oosyyyyyhhhhhhyhhhhyyyyyy:..............................o
o'''''''''''.....:. +ss:-//://:://+ooooo++//++osyyyhhhhyyyhhhhyyssy+...............................o
o'''''''''''.'.:/.'.hso-..::::/:::/+++oo+++/+++oo++oosssyyyhyyyssy+................................o
o'''''''''''.:+/-''/hs:-...-:::::::///++oooooooossssssyyyyyyyysss/.................................o
o'''''''''-:::::.''/yo:--....---:::://+++oossyyyyhhhhhhhyyyyysso:'''...............................o
o''''..-:/:----.''':ss/:---....----::///++oosyhhhhhhhhhyyyyysoo.''''...''''........................o
o.----.--.''...''''-+ss/:::---...----:://++osyyhyyyhhhhysyssoyy..'''''.......''''''''..............o
s--.'''''''''''''''./+os+:::::---------::/+oossyyyyyyyysssoo+ms...'''''.........''''''''''''''.....o
o''''''''''''''''''':+ooss+//:::::-------:/+oossssssssssoos+ymo...''''''...........''''''''''''''''o
o'''''''''''''''''''.+oossyso////::::::::::////++++++oooss+ymdo....''''''.................'''''''''o
o''''''''''''''''''''-osssyyhys++////://///+++oooooosssyssmNmho....'''''''................'....''''o
s:::::::::::::::::::::+yyhhhhdddyssoooooooosssyyyyyyyhyymNNNmdy////::::::::::::::::::::::::::::::::s
`
)

func getC(c []byte) byte {
	return c[rand.Intn(len(c))]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Ctrl struct {
	b []byte
	c []byte
	p []int
}

var args struct {
	Time time.Duration `arg:"-t,help:speed of animation in ms (Default: 25)"`
	//Time time.Duration `arg:"-t,required,help:speed of animation in ms"`
	Mode int `arg:"-m,help:1 positive - 2 negative (Default: 2)"`
}

var source string

func main() {
	tw.Cls()
	arg.MustParse(&args)
	if args.Time == 0 {
		args.Time = 25
	}
	if args.Mode == 1 {
		source = tw2
	} else if args.Mode == 2 {
		source = tw1
	} else {
		source = tw2
	}

	r := bytes.NewBuffer([]byte(source))

	var err error
	k := 0

	img := make([]byte, 0)
	m := make([]Ctrl, 0)

	for {
		if img, err = r.ReadBytes('\n'); err == io.EOF {
			break
		}
		img = img[:len(img)-1]
		m = append(m, Ctrl{
			b: img,
			c: make([]byte, len(img)),
			p: rand.Perm(len(img)), //define which position will be set in loop
		})
		k++
	}
	nLines := k - 1
	//Debug("Lines: %d, %s", nLines, m[nLines].b)

	lLength := len(m[1].b)
	result := make([][]byte, 30)
	for i := range result {
		result[i] = make([]byte, lLength)
	}

	writer := uilive.New()

	// start listening for updates and render
	writer.Start()

	for i := 0; i <= lLength; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		// loop over picture
		for k := 0; k < nLines; k++ {
			// loop over line
			for j := 0; j < lLength; j++ {
				// if control bit alread set continue
				if m[k].c[j] == 1 {
					continue
				}
				// if location is value of permutation set value and control bit
				if j == m[k].p[i] {
					result[k][j] = m[k].b[j]
					m[k].c[j] = 1
				} else {
					result[k][j] = getC([]byte(letterBytes))
				}
			}
			fmt.Fprintf(writer, "%s\n", result[k])
		}
		time.Sleep(time.Millisecond * args.Time)
	}

	writer.Stop() // flush and stop rendering
}
