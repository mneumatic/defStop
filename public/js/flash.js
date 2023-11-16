const flash = document.getElementById('flash-msg')
if (flash) {
  setTimeout(() => {
    flash.style.opacity = '0'
    flash.style.pointerEvents = 'none'
  }, 2000)
}
