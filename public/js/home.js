const span = document.querySelector('.header-span')
const quote = document.getElementById('quote')
const mottos = [
  'defStop aims to make sure your next dump is not at a dump.',
  'defStop aims to make all the small logs on the road a smoother ride.',
  'Honey?! Did you pack any toilet paper?'
]

const items = [
  'toilet paper.',
  'semi parking.',
  'a dog park.',
  'a shower.',
  'a vending machine.'
]

function generateArrayElement (arr) {
  return arr[Math.floor(Math.random() * arr.length)]
}

span.innerText = generateArrayElement(items)
quote.innerText = generateArrayElement(mottos)
