// 检查是否登录
export function checkoutLogin () {
  var isLogin = localStorage.getItem("_USERNAME_")
  return isLogin ? true : false
}

// 保存用户名到localstorage
export function setUsername (username) {
  localStorage.setItem('_USERNAME_', username)
}

// 删除用户名从缓存
export function removeUsername () {
  localStorage.removeItem("_USERNAME_")
}