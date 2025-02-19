// Copyright © 2018 Kent Gibson <warthog618@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package charset

var (
	hindiDecoder = Decoder{
		0x00: '\u0981',
		0x01: '\u0982',
		0x02: '\u0983',
		0x03: 'अ',
		0x04: 'आ',
		0x05: 'इ',
		0x06: 'ई',
		0x07: 'उ',
		0x08: 'ऊ',
		0x09: 'ऋ',
		0x0a: '\n',
		0x0b: 'ऌ',
		0x0c: 'ऍ',
		0x0d: '\r',
		0x0e: 'ऎ',
		0x0f: 'ए',
		0x10: 'ऐ',
		0x11: 'ऑ',
		0x12: 'ऒ',
		0x13: 'ओ',
		0x14: 'औ',
		0x15: 'क',
		0x16: 'ख',
		0x17: 'ग',
		0x18: 'घ',
		0x19: 'ङ',
		0x1a: 'च',
		0x1b: 0x1b,
		0x1c: 'छ',
		0x1d: 'ज',
		0x1e: 'झ',
		0x1f: 'ञ',
		0x20: 0x20,
		0x21: '!',
		0x22: 'ट',
		0x23: 'ठ',
		0x24: 'ड',
		0x25: 'ढ',
		0x26: 'ण',
		0x27: 'त',
		0x28: ')',
		0x29: '(',
		0x2a: 'थ',
		0x2b: 'द',
		0x2c: ',',
		0x2d: 'ध',
		0x2e: '.',
		0x2f: 'न',
		0x30: '0',
		0x31: '1',
		0x32: '2',
		0x33: '3',
		0x34: '4',
		0x35: '5',
		0x36: '6',
		0x37: '7',
		0x38: '8',
		0x39: '9',
		0x3a: ':',
		0x3b: ';',
		0x3c: 'ऩ',
		0x3d: 'प',
		0x3e: 'फ',
		0x3f: '?',
		0x40: 'ब',
		0x41: 'भ',
		0x42: 'म',
		0x43: 'य',
		0x44: 'र',
		0x45: 'ऱ',
		0x46: 'ल',
		0x47: 'ळ',
		0x48: 'ऴ',
		0x49: 'व',
		0x4a: 'श',
		0x4b: 'ष',
		0x4c: 'स',
		0x4d: 'ह',
		0x4e: '\u093c',
		0x4f: 'ऽ',
		0x50: '\u093e',
		0x51: '\u093f',
		0x52: '\u0940',
		0x53: '\u0941',
		0x54: '\u0942',
		0x55: '\u0943',
		0x56: '\u0944',
		0x57: '\u0945',
		0x58: '\u0946',
		0x59: '\u0947',
		0x5a: '\u0948',
		0x5b: '\u0949',
		0x5c: '\u094a',
		0x5d: '\u094b',
		0x5e: '\u094c',
		0x5f: '\u094d',
		0x60: 'ॐ',
		0x61: 'a',
		0x62: 'b',
		0x63: 'c',
		0x64: 'd',
		0x65: 'e',
		0x66: 'f',
		0x67: 'g',
		0x68: 'h',
		0x69: 'i',
		0x6a: 'j',
		0x6b: 'k',
		0x6c: 'l',
		0x6d: 'm',
		0x6e: 'n',
		0x6f: 'o',
		0x70: 'p',
		0x71: 'q',
		0x72: 'r',
		0x73: 's',
		0x74: 't',
		0x75: 'u',
		0x76: 'v',
		0x77: 'w',
		0x78: 'x',
		0x79: 'y',
		0x7a: 'z',
		0x7b: 'ॲ',
		0x7c: 'ॻ',
		0x7d: 'ॼ',
		0x7e: 'ॾ',
		0x7f: 'ॿ',
	}
	hindiExtDecoder = Decoder{
		0x00: '@',
		0x01: '£',
		0x02: '$',
		0x03: '¥',
		0x04: '¿',
		0x05: '"',
		0x06: '¤',
		0x07: '%',
		0x08: '&',
		0x09: '\'',
		0x0a: '\f',
		0x0b: '*',
		0x0c: '+',
		0x0d: '\r',
		0x0e: '-',
		0x0f: '/',
		0x10: '<',
		0x11: '=',
		0x12: '>',
		0x13: '¡',
		0x14: '^',
		0x15: '¡',
		0x16: '_',
		0x17: '#',
		0x18: '*',
		0x19: '।',
		0x1a: '॥',
		0x1b: 0x1b,
		0x1c: '०',
		0x1d: '१',
		0x1e: '२',
		0x1f: '३',
		0x20: '४',
		0x21: '५',
		0x22: '६',
		0x23: '७',
		0x24: '८',
		0x25: '९',
		0x26: '\u0951',
		0x27: '\u0952',
		0x28: '{',
		0x29: '}',
		0x2a: '\u0953',
		0x2b: '\u0954',
		0x2c: 'क़',
		0x2d: 'ख़',
		0x2e: 'ग़',
		0x2f: '\\',
		0x30: 'ज़',
		0x31: 'ड़',
		0x32: 'ढ़',
		0x33: 'फ़',
		0x34: 'य़',
		0x35: 'ॠ',
		0x36: 'ॡ',
		0x37: '\u0962',
		0x38: '\u0963',
		0x39: '॰',
		0x3a: 'ॱ',
		0x3c: '[',
		0x3d: '~',
		0x3e: ']',
		0x40: '|',
		0x41: 'A',
		0x42: 'B',
		0x43: 'C',
		0x44: 'D',
		0x45: 'E',
		0x46: 'F',
		0x47: 'G',
		0x48: 'H',
		0x49: 'I',
		0x4a: 'J',
		0x4b: 'K',
		0x4c: 'L',
		0x4d: 'M',
		0x4e: 'N',
		0x4f: 'O',
		0x50: 'P',
		0x51: 'Q',
		0x52: 'R',
		0x53: 'S',
		0x54: 'T',
		0x55: 'U',
		0x56: 'V',
		0x57: 'W',
		0x58: 'X',
		0x59: 'Y',
		0x5a: 'Z',
		0x65: '€',
	}
	hindiEncoder    Encoder
	hindiExtEncoder Encoder
)

func generateHindiEncoder() Encoder {
	return generateEncoder(hindiDecoder)
}

func generateHindiExtEncoder() Encoder {
	return generateEncoder(hindiExtDecoder)
}
