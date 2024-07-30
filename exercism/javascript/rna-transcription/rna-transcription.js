//
// This is only a SKELETON file for the 'RNA Transcription' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

const complements = {
  G: 'C',
  C: 'G',
  T: 'A',
  A: 'U'
}

export const toRna = rna =>
  !rna ? '' : rna.split('').map(n => complements[n]).join('')
