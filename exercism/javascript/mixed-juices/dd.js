import { timeToMixJuice, limesToCut, remainingOrders } from './mixed-juices'

const test = (name, received, expected) =>
  JSON.stringify(received) === JSON.stringify(expected)
    ? console.log(name, 'pass')
    : console.error(name, 'fail', { expected, received })

test(
  'Test 3-1',
  remainingOrders(7, [
    'Tropical Island',
    'Energizer',
    'Limetime',
    'All or Nothing',
    'Pure Strawberry Joy'
  ]),
  ['All or Nothing', 'Pure Strawberry Joy']
)

test(
  'Test 3-2',
  remainingOrders(13, [
    'Pure Strawberry Joy',
    'Pure Strawberry Joy',
    'Vitality',
    'Tropical Island',
    'All or Nothing',
    'All or Nothing',
    'All or Nothing',
    'Green Garden',
    'Limetime'
  ]),
  ['All or Nothing', 'Green Garden', 'Limetime']
)

test(
  'Test 3-3',
  remainingOrders(12, [
    'Energizer',
    'Green Garden',
    'Ruby Glow',
    'Pure Strawberry Joy',
    'Tropical Island',
    'Limetime'
  ]),
  []
)

test(
  'Test 3-4',
  remainingOrders(0.2, ['Bananas Gone Wild', 'Pure Strawberry Joy']),
  ['Pure Strawberry Joy']
)
