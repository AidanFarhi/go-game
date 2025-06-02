const grid = document.getElementById("game-grid")

const populateGrid = async () => {
    const response = await fetch("/gamegrid")
    const data = await response.text()
    grid.innerHTML = data
}

populateGrid()
