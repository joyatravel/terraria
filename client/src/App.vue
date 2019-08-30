<template>
  <div id="app">
    <router-view />
  </div>
</template>

<style lang="scss">
@import "@/styles/fonts.scss";

body {
  font-family: $sans;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

// prettier-ignore
h1, h2, h3, h4, h5, h6, p { margin: 0; }

// Flexbox shortcuts:
.flex {
  display: flex;
  flex-direction: column;
}
.center-children {
  align-items: center;
  justify-content: center;
}

// Font overrides:
.sans {
  font-family: $sans;
}
.mono {
  font-family: $mono;
}
</style>
