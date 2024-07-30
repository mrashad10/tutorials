/// <reference path="./global.d.ts" />
// @ts-check

export const cookingStatus = time => {
  if (time === undefined) return 'You forgot to set the timer.'
  return time === 0 ? 'Lasagna is done.' : 'Not done, please wait.'
}

export const preparationTime = (layers, average = 2) => layers.length * average

export const quantities = layers => ({
  noodles: layers.filter(element => element === 'noodles').length * 50,
  sauce: layers.filter(element => element === 'sauce').length * 0.2
})

export const addSecretIngredient = (friendsList, myList) => {
  myList.push(friendsList[friendsList.length - 1])
}

export const scaleRecipe = (recipe, count = 2) => {
  const result = { ...recipe }
  if (count === 2) return result

  for (const key in result) {
    result[key] *= count / 2
  }

  return result
}
