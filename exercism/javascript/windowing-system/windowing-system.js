// @ts-check

/**
 * Implement the classes etc. that are needed to solve the
 * exercise in this file. Do not forget to export the entities
 * you defined so they are available for the tests.
 */

function Size(width = 80, height = 60) {
    this.width = width
    this.height = height

}

Size.prototype.resize = function (width, height) {
    this.width = Math.max(1, width)
    this.height = Math.max(1, height)
}

function Position(x = 0, y = 0) {
    this.x = Math.max(0, x)
    this.y = Math.max(0, y)
}

Position.prototype.move = function (x, y) {
    this.x = x
    this.y = y
}

class ProgramWindow {
    constructor() {
        this.screenSize = new Size(800, 600)
        this.size = new Size()
        this.position = new Position()
    }

    resize(size) {
        const maxAcceptableWidth = this.screenSize.width - this.position.x
        const maxAcceptableHeight = this.screenSize.height - this.position.y

        const newWidth = Math.min(maxAcceptableWidth, size.width)
        const newHeight = Math.min(maxAcceptableHeight, size.height)

        this.size.resize(newWidth, newHeight)
    }

    move(position) {
        const { width, height } = this.screenSize
        const { width: sizeWidth, height: sizeHeight } = this.size

        const maxAcceptableX = width - sizeWidth
        const maxAcceptableY = height - sizeHeight

        const newX = Math.min(maxAcceptableX, position.x)
        const newY = Math.min(maxAcceptableY, position.y)

        this.position.move(newX, newY)
    }
}

function changeWindow(programWindow) {
    const size = new Size(400, 300)
    const position = new Position(100, 150)

    programWindow.resize(size)
    programWindow.move(position)

    return programWindow
}

module.exports = {
    Size,
    Position,
    ProgramWindow,
    changeWindow,
}