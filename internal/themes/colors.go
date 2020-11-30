package themes

var AllThemes = make(map[string]string)
var ThemesList []string

func init() {
	AllThemes = map[string]string{
		"After Glow":          AfterGlow,
		"Argonaut":            Argonaut,
		"Ayu Dark":            AyuDark,
		"Ayu Mirage":          AyuMirage,
		"Base16 Default Dark": Base16DefaultDark,
		"Blood Moon":          BloodMoon,
		"Breeze":              Breeze,
		"Campbell":            Campbell,
		"Challenger Deep":     ChallengerDeep,
		"Cobalt2":             Cobalt2,
		"Cyber Punk Neon":     CyberPunkNeon,
		"Darcula":             Darcula,
		"Default Theme":       DefaultTheme,
		"Doom One":            DoomOne,
		"Dracula":             Dracula,
		"Falcon":              Falcon,
		"Flat Remix":          FlatRemix,
		"Gotham":              Gotham,
		"Solarized Light":     SolarizedLight,
	}

	for theme := range AllThemes {
		ThemesList = append(ThemesList, theme)
	}
}

const AfterGlow string = `
colors:

  primary:
    background: '#2c2c2c'
    foreground: '#d6d6d6'

    dim_foreground:    '#dbdbdb'
    bright_foreground: '#d9d9d9'
    dim_background:    '#202020'
    bright_background: '#3a3a3a'

  cursor:
    text:   '#2c2c2c'
    cursor: '#d9d9d9'

  normal:
    black:   '#1c1c1c'
    red:     '#bc5653'
    green:   '#909d63'
    yellow:  '#ebc17a'
    blue:    '#7eaac7'
    magenta: '#aa6292'
    cyan:    '#86d3ce'
    white:   '#cacaca'

  bright:
    black:   '#636363'
    red:     '#bc5653'
    green:   '#909d63'
    yellow:  '#ebc17a'
    blue:    '#7eaac7'
    magenta: '#aa6292'
    cyan:    '#86d3ce'
    white:   '#f7f7f7'

  dim:
    black:   '#232323'
    red:     '#74423f'
    green:   '#5e6547'
    yellow:  '#8b7653'
    blue:    '#556b79'
    magenta: '#6e4962'
    cyan:    '#5c8482'
    white:   '#828282'
`

const Argonaut string = `
colors:

  primary:
    background: '#292C3E'
    foreground: '#EBEBEB'

  cursor:
   text: '#FF261E'
   cursor: '#FF261E'

  normal:
    black:   '#0d0d0d'
    red:     '#FF301B'
    green:   '#A0E521'
    yellow:  '#FFC620'
    blue:    '#1BA6FA'
    magenta: '#8763B8'
    cyan:    '#21DEEF'
    white:   '#EBEBEB'

  bright:
    black:   '#6D7070'
    red:     '#FF4352'
    green:   '#B8E466'
    yellow:  '#FFD750'
    blue:    '#1BA6FA'
    magenta: '#A578EA'
    cyan:    '#73FBF1'
    white:   '#FEFEF8'
`

const AyuMirage string = `
colors:

  primary:
    background: '#202734'
    foreground: '#CBCCC6'

  normal:
    black: '#191E2A'
    red: '#FF3333'
    green: '#BAE67E'
    yellow: '#FFA759'
    blue: '#73D0FF'
    magenta: '#FFD580'
    cyan: '#95E6CB'
    white: '#C7C7C7'

  bright:
    black: '#686868'
    red: '#F27983'
    green: '#A6CC70'
    yellow: '#FFCC66'
    blue: '#5CCFE6'
    magenta: '#FFEE99'
    cyan: '#95E6CB'
    white: '#FFFFFF'
`

const Base16DefaultDark string = `
colors:

  primary:
    background: '#181818'
    foreground: '#d8d8d8'

  cursor:
    text: '#d8d8d8'
    cursor: '#d8d8d8'

  normal:
    black:   '#181818'
    red:     '#ab4642'
    green:   '#a1b56c'
    yellow:  '#f7ca88'
    blue:    '#7cafc2'
    magenta: '#ba8baf'
    cyan:    '#86c1b9'
    white:   '#d8d8d8'

  bright:
    black:   '#585858'
    red:     '#ab4642'
    green:   '#a1b56c'
    yellow:  '#f7ca88'
    blue:    '#7cafc2'
    magenta: '#ba8baf'
    cyan:    '#86c1b9'
    white:   '#f8f8f8'
`

const BloodMoon string = `
colors:

  primary:
    background: '#10100E'
    foreground: '#C6C6C4'

  normal:
    black:   '#10100E'
    red:     '#C40233'
    green:   '#009F6B'
    yellow:  '#FFD700'
    blue:    '#0087BD'
    magenta: '#9A4EAE'
    cyan:    '#20B2AA'
    white:   '#C6C6C4'

  bright:
    black:   '#696969'
    red:     '#FF2400'
    green:   '#03C03C'
    yellow:  '#FDFF00'
    blue:    '#007FFF'
    magenta: '#FF1493'
    cyan:    '#00CCCC'
    white:   '#FFFAFA'
`

const Breeze string = `
colors:

  primary:
    background: '0x232627'
    foreground: '0xfcfcfc'

    dim_foreground: '0xeff0f1'
    bright_foreground: '0xffffff'
    dim_background: '0x31363b'
    bright_background: '0x000000'

  normal:
    black: '0x232627'
    red: '0xed1515'
    green: '0x11d116'
    yellow: '0xf67400'
    blue: '0x1d99f3'
    magenta: '0x9b59b6'
    cyan: '0x1abc9c'
    white: '0xfcfcfc'

  bright:
    black: '0x7f8c8d'
    red: '0xc0392b'
    green: '0x1cdc9a'
    yellow: '0xfdbc4b'
    blue: '0x3daee9'
    magenta: '0x8e44ad'
    cyan: '0x16a085'
    white: '0xffffff'

  dim:
    black: '0x31363b'
    red: '0x783228'
    green: '0x17a262'
    yellow: '0xb65619'
    blue: '0x1b668f'
    magenta: '0x614a73'
    cyan: '0x186c60'
    white: '0x63686d'
`

const Campbell string = `
colors:

  primary:
    background: '0x0c0c0c'
    foreground: '0xcccccc'

  normal:
    black:      '0x0c0c0c'
    red:        '0xc50f1f'
    green:      '0x13a10e'
    yellow:     '0xc19c00'
    blue:       '0x0037da'
    magenta:    '0x881798'
    cyan:       '0x3a96dd'
    white:      '0xcccccc'

  bright:
    black:      '0x767676'
    red:        '0xe74856'
    green:      '0x16c60c'
    yellow:     '0xf9f1a5'
    blue:       '0x3b78ff'
    magenta:    '0xb4009e'
    cyan:       '0x61d6d6'
    white:      '0xf2f2f2'
`

const ChallengerDeep string = `
colors:

  primary:
    background: '0x1e1c31'
    foreground: '0xcbe1e7'

  cursor:
    text: '0xff271d'
    cursor: '0xfbfcfc'

  normal:
    black:   '0x141228'
    red:     '0xff5458'
    green:   '0x62d196'
    yellow:  '0xffb378'
    blue:    '0x65b2ff'
    magenta: '0x906cff'
    cyan:    '0x63f2f1'
    white:   '0xa6b3cc'

  bright:
    black:   '0x565575'
    red:     '0xff8080'
    green:   '0x95ffa4'
    yellow:  '0xffe9aa'
    blue:    '0x91ddff'
    magenta: '0xc991e1'
    cyan:    '0xaaffe4'
    white:   '0xcbe3e7'
`

const Cobalt2 string = `
colors:

  primary:
    background: '0x122637'
    foreground: '0xffffff'


  cursor:
    text: '0x122637'
    cursor: '0xf0cb09'


  normal:
    black:   '0x000000'
    red:     '0xff0000'
    green:   '0x37dd21'
    yellow:  '0xfee409'
    blue:    '0x1460d2'
    magenta: '0xff005d'
    cyan:    '0x00bbbb'
    white:   '0xbbbbbb'

  bright:
    black:   '0x545454'
    red:     '0xf40d17'
    green:   '0x3bcf1d'
    yellow:  '0xecc809'
    blue:    '0x5555ff'
    magenta: '0xff55ff'
    cyan:    '0x6ae3f9'
    white:   '0xffffff'
`

const CyberPunkNeon string = `
colors:

  primary:
    background: "0x000b1e"
    foreground: "0x0abdc6"

  cursor:
    text:   "0x000b1e"
    cursor: "0x0abdc6"

  normal:
    black:   "0x123e7c"
    red:     "0xff0000"
    green:   "0xd300c4"
    yellow:  "0xf57800"
    blue:    "0x123e7c"
    magenta: "0x711c91"
    cyan:    "0x0abdc6"
    white:   "0xd7d7d5"

  bright:
    black:   "0x1c61c2"
    red:     "0xff0000"
    green:   "0xd300c4"
    yellow:  "0xf57800"
    blue:    "0x00ff00"
    magenta: "0x711c91"
    cyan:    "0x0abdc6"
    white:   "0xd7d7d5"
`

const Darcula string = `
colors:

  primary:
    background: '0x282a36'
    foreground: '0xf8f8f2'

  normal:
    black:   '0x000000'
    red:     '0xff5555'
    green:   '0x50fa7b'
    yellow:  '0xf1fa8c'
    blue:    '0xcaa9fa'
    magenta: '0xff79c6'
    cyan:    '0x8be9fd'
    white:   '0xbfbfbf'

  bright:
    black:   '0x282a35'
    red:     '0xff6e67'
    green:   '0x5af78e'
    yellow:  '0xf4f99d'
    blue:    '0xcaa9fa'
    magenta: '0xff92d0'
    cyan:    '0x9aedfe'
    white:   '0xe6e6e6'
`

const DefaultTheme string = `
colors:

  primary:
    background: '0x1d1f21'
    foreground: '0xeaeaea'

  normal:
    black:   '0x000000'
    red:     '0xd54e53'
    green:   '0xb9ca4a'
    yellow:  '0xe6c547'
    blue:    '0x7aa6da'
    magenta: '0xc397d8'
    cyan:    '0x70c0ba'
    white:   '0xffffff'

  bright:
    black:   '0x666666'
    red:     '0xff3334'
    green:   '0x9ec400'
    yellow:  '0xe7c547'
    blue:    '0x7aa6da'
    magenta: '0xb77ee0'
    cyan:    '0x54ced6'
    white:   '0xffffff'

  dim:
    black:   '0x333333'
    red:     '0xf2777a'
    green:   '0x99cc99'
    yellow:  '0xffcc66'
    blue:    '0x6699cc'
    magenta: '0xcc99cc'
    cyan:    '0x66cccc'
    white:   '0xdddddd'
`

const Dracula string = `
colors:

  primary:
    background: '0x282a36'
    foreground: '0xf8f8f2'
 
  normal:
    black:   '0x000000'
    red:     '0xff5555'
    green:   '0x50fa7b'
    yellow:  '0xf1fa8c'
    blue:    '0xbd93f9'
    magenta: '0xff79c6'
    cyan:    '0x8be9fd'
    white:   '0xbbbbbb'
 
  bright:
    black:   '0x555555'
    red:     '0xff5555'
    green:   '0x50fa7b'
    yellow:  '0xf1fa8c'
    blue:    '0xcaa9fa'
    magenta: '0xff79c6'
    cyan:    '0x8be9fd'
    white:   '0xffffff'
`

const Falcon string = `
colors:

  primary:
    background: '0x020221'
    foreground: '0xb4b4b9'

  cursor:
    text: '0x020221'
    cursor: '0xffe8c0'

  normal:
    black:   '0x000004'
    red:     '0xff3600'
    green:   '0x718e3f'
    yellow:  '0xffc552'
    blue:    '0x635196'
    magenta: '0xff761a'
    cyan:    '0x34bfa4'
    white:   '0xb4b4b9'

  bright:
    black:   '0x020221'
    red:     '0xff8e78'
    green:   '0xb1bf75'
    yellow:  '0xffd392'
    blue:    '0x99a4bc'
    magenta: '0xffb07b'
    cyan:    '0x8bccbf'
    white:   '0xf8f8ff'
`

const FlatRemix string = `
colors:
  primary:
    background: '0x272a34'
    foreground: '0xFFFFFF'

  normal:
    black:   '0x1F2229'
    red:     '0xEC0101'
    green:   '0x47D4B9'
    yellow:  '0xFF8A18'
    blue:    '0x277FFF'
    magenta: '0xD71655'
    cyan:    '0x05A1F7'
    white:   '0xFFFFFF'


  bright:
    black:   '0x1F2229'
    red:     '0xD41919'
    green:   '0x5EBDAB'
    yellow:  '0xFEA44C'
    blue:    '0x367bf0'
    magenta: '0xBF2E5D'
    cyan:    '0x49AEE6'
    white:   '0xFFFFFF'
`

const Gotham string = `
colors:

  primary:
    background: '0x0a0f14'
    foreground: '0x98d1ce'

  normal:
    black: '0x0a0f14'
    red: '0xc33027'
    green: '0x26a98b'
    yellow: '0xedb54b'
    blue: '0x195465'
    magenta: '0x4e5165'
    cyan: '0x33859d'
    white: '0x98d1ce'

  bright:
    black: '0x10151b'
    red: '0xd26939'
    green: '0x081f2d'
    yellow: '0x245361'
    blue: '0x093748'
    magenta: '0x888ba5'
    cyan: '0x599caa'
    white: '0xd3ebe9'
`

const DoomOne string = `
colors:

  primary:
    background: '0x282c34'
    foreground: '0xbbc2cf'

  normal:
    black:   '0x282c34'
    red:     '0xff6c6b'
    green:   '0x98be65'
    yellow:  '0xecbe7b'
    blue:    '0x51afef'
    magenta: '0xc678dd'
    cyan:    '0x46d9ff'
    white:   '0xbbc2cf'
`

const SolarizedLight string = `
colors:

  primary:
    background: '#fdf6e3'
    foreground: '#657b83'

  normal:
    black:   '#073642'
    red:     '#dc322f'
    green:   '#859900'
    yellow:  '#b58900'
    blue:    '#268bd2'
    magenta: '#d33682'
    cyan:    '#2aa198'
    white:   '#eee8d5'

  bright:
    black:   '#002b36'
    red:     '#cb4b16'
    green:   '#586e75'
    yellow:  '#657b83'
    blue:    '#839496'
    magenta: '#6c71c4'
    cyan:    '#93a1a1'
    white:   '#fdf6e3'
`

const AyuDark string = `
colors:

  primary:
    background: '#0A0E14'
    foreground: '#B3B1AD'

  normal:
    black: '#01060E'
    red: '#EA6C73'
    green: '#91B362'
    yellow: '#F9AF4F'
    blue: '#53BDFA'
    magenta: '#FAE994'
    cyan: '#90E1C6'
    white: '#C7C7C7'

  bright:
    black: '#686868'
    red: '#F07178'
    green: '#C2D94C'
    yellow: '#FFB454'
    blue: '#59C2FF'
    magenta: '#FFEE99'
    cyan: '#95E6CB'
    white: '#FFFFFF'
`
