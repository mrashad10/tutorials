import { newIsPangram, oldIsPangram } from './pangram/pangram'

// Example sentence to test
const sentences = [
  '',
  'not pangram',
  'The quick brown fox jumps over the lazy dog',
  'Bright vixens jump; dozy fowl quack',
  'Sphinx of black quartz, judge my vow',
  'The five boxing wizards jump quickly',
  'Pack my box with five dozen liquor jugs',
  'Jinxed wizards pluck ivy from the big quilt',
  'The jay, pig, fox, zebra, and my wolves quack',
  'Blowzy night-frumps vex’d Jack Q',
  'Quick zephyrs blow, vexing daft Jim',
  'Waltz, nymph, for quick jigs vex Bud',
  'abcdefghijklmnopqrstuvwxyz',
  'the quick brown fox jumps over the lazy dog',
  'five boxing wizards jump quickly at it',
  'the_quick_brown_fox_jumps_over_the_lazy_dog',
  'the 1 quick brown fox jumps over the 2 lazy dogs',
  '7h3 qu1ck brown fox jumps ov3r 7h3 lazy dog',
  '"Five quacking Zephyrs jolt my wax bed."',
  'the quick brown fox jumps over with lazy FX'
]

/**
 * Benchmark the performance of two pangram functions.
 *
 * @param {number} loopsCount - The number of times to loop the benchmark.
 * @return {undefined} No return value.
 */
const benchmarkPangramFunctions = (loopsCount = 1) => {
  let newDuration = 0,
    oldDuration = 0

  // Benchmark for isPangram
  newDuration = benchmark(newIsPangram, loopsCount)

  // Benchmark for oldIsPangram
  oldDuration = benchmark(oldIsPangram, loopsCount)

  // Display the difference in time
  const diffTime = Math.abs(newDuration - oldDuration)
  const newTimeOverOldPercentage = (newDuration / oldDuration * 100).toFixed(2)
  const oldTimeOverNewPercentage = (oldDuration / newDuration * 100).toFixed(2)

  console.log(
    `For ${sentences.length * loopsCount} sentences,
      New isPangram: ${newDuration} ms.
      Old isPangram: ${oldDuration} ms.
      Differance in time: ${diffTime} ms.
      New time over old: ${newTimeOverOldPercentage}%
      Enhancment in performance: ${oldTimeOverNewPercentage}%`
  )
}

/**
 * Measures the performance of a given function by executing it a specified number of times.
 *
 * @param {function} func - The function to be benchmarked.
 * @param {number} [loopsCount=1] - The number of times the function should be executed.
 * @return {number} - The duration of the benchmark in milliseconds.
 */
const benchmark = (func, loopsCount = 1) => {
  let startTime, endTime, duration

  startTime = performance.now()
  for (let index = 1; index <= loopsCount; index++) {
    sentences.forEach(sentence => func(sentence))
  }
  endTime = performance.now()

  duration = endTime - startTime
  return duration
}

// Run the benchmark
const loopsCount = process.argv[2] ? parseInt(process.argv[2]) : 1000
benchmarkPangramFunctions(loopsCount)
