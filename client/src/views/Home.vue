<template>
  <div class="home flex">
    <header class="flex center-children">
      <div class="inverse">
        <h1 class="title">Terraria</h1>
      </div>
      <div class="spacer" />
      <h1 class="tag">@tugolo</h1>
    </header>
    <main>
      <section class="connect">
        <h2 class="title">Connect</h2>
        <div
          class="row"
          v-for="[key, val] in Object.entries(server)"
          :key="key"
        >
          <h3 class="key">{{ key }}</h3>
          <p class="val mono">{{ val }}</p>
        </div>
      </section>

      <section class="status">
        <h2 class="title">Status</h2>
        <span v-if="!status.error">
          <div class="row">
            <h3 class="key">world</h3>
            <p class="val mono">{{ status.world }}</p>
          </div>
          <div class="row">
            <h3 class="key">players online</h3>
            <p class="val mono">{{ status.playerCount }}</p>
          </div>
          <div class="row">
            <h3 class="key">player capacity</h3>
            <p class="val mono">{{ status.maxPlayerCount }}</p>
          </div>
        </span>
        <span v-else>
          <p class="error">The server seems to be offline.</p>
        </span>
      </section>

      <section class="players" v-if="!isEmpty(status.players)">
        <h2 class="title">Players Online</h2>
        <div class="tags">
          <div class="tag" v-for="name in status.players" :key="name">
            <p>{{ name }}</p>
          </div>
        </div>
      </section>
    </main>

    <div class="download" @click="download">
      <button>Download World File</button>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import terraria from "@/services/terraria";
import isEmpty from "lodash/isEmpty";
// import HelloWorld from "@/components/HelloWorld.vue";

export default {
  data: () => ({
    status: {
      error: null,
    },
    server: {
      host: "terraria.tugolo.com",
      port: 6969,
    },
  }),
  created() {
    terraria
      .getStatus()
      .then(status => (this.status = status), err => (this.status.error = err));
  },
  methods: {
    isEmpty,
    download() {
      terraria.downloadWorld();
    },
  },
  components: {
    // HelloWorld
  },
};
</script>

<style lang="scss" scoped>
@import "@/styles/breakpoints.scss";

.home {
  align-items: center;
  padding-top: 75px;

  @include breakpoint(mobileonly) {
    padding-top: 35px;
  }
}

header {
  .title {
    font-size: 40pt;
    font-weight: 700;
  }

  .tag {
    margin-top: 12px;
    font-size: 18pt;
    font-weight: 600;
  }

  // prettier-ignore
  .spacer { width: 6px; }

  .inverse {
    padding: 12px;
    color: white;
    background: black;
  }

  // prettier-ignore
  @include breakpoint(mobileonly) {
    .title { font-size: 28pt; }
    .tag { font-size: 16pt; }
  }
}

main {
  margin-top: 32px;
  width: 300px;

  section {
    margin-bottom: 22px;

    // prettier-ignore
    .title { margin-bottom: 5px; }

    .row {
      display: flex;
      align-items: center;

      // prettier-ignore
      .key { font-weight: 600; }

      .val {
        margin-left: 14px;
        font-size: 14pt;
        color: rgb(60, 60, 60);
      }
    }
  }

  .status {
    // prettier-ignore
    .row .key { width: 150px; }
    .error {
      font-size: 13pt;
      font-weight: 500;
      color: rgb(212, 0, 0);
    }
  }

  .players .tags {
    display: flex;
    flex-wrap: wrap;

    .tag {
      margin: 2px;
      padding: 6px 8px;
      border: solid rgb(70, 70, 70) 2px;
      border-radius: 2px;
      box-sizing: border-box;
      font-weight: 500;
    }
  }

  @include breakpoint(mobileonly) {
    margin-top: 20px;
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;

    // prettier-ignore
    section {
      .title { font-size: 17pt; }
      .row {
        .key { font-size: 13.5pt; }
        .val { font-size: 13.5pt; }
      }
    }
  }
}

.download {
  margin: 20px 0 30px 0;

  button {
    padding: 8px 10px;
    border: none;
    border-radius: 2px;
    outline: none;

    cursor: pointer;
    font-weight: 500;
    color: white;
    background: black;

    // prettier-ignore
    &:hover { background: rgb(50, 50, 50); }
    transition: background 50ms ease-in-out;
  }

  // prettier-ignore
  @include breakpoint(mobileonly) { margin-top: 15px; }
}
</style>
