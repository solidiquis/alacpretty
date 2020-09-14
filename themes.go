package main

const afterGlow string = `
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

const argonaut string = `
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

const ayuMirage string = `
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

const base16DefaultDark string = `
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

const bloodMoon string = `
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

const defaultTheme string = `
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

const solarizedLight string = `
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

const ayuDark string = `
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
