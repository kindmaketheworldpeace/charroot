<template>
  <div class="back-ground" :style="{backgroundImage:`url(${img})`}">
    <div class="login">
      <div class="account">
        <span class="account-header"> 用户名:</span>
        <Input  v-model="username" placeholder="请输入用户名" class="account-input" />
      </div>
      <div class="commit">
        <i-button   @click="entry" type="success">进入</i-button>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Home',
    data() {
      return {
        img: require('../assets/background.jpg'),
        username:""

      }
    },
    methods: {
      entry: function () {
        let  params= {
          username:this.username
        }
        this.$store.dispatch('user/entry', params).then(res => {
            if (res.status===200) {
              this.$router.push({path: '/chatroom'})
            } else {
              this.$message.error("进入失败！")
            }
          }
        )


      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .login {
    position: absolute;
    border-radius: 3px;
    width: 400px;
    height: 300px;
    top: calc(50% - 150px);
    left: calc(50% - 200px);
    border-color: #dedede;
    border-style: solid;
    border-width: 2px;
    background: rgba(147, 147, 147, 0.5);

  }

  .header {
    height: 60px;
    background-color: #dedede;

  }

  .back-ground {

    height: 100%;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-attachment: fixed;
  }

  .account {
    position: absolute;
    top: 40%;
    left: 10%;
    width: 80% !important;
  }

  .account-input {

    width: 70%;
    margin-left: 10px;
    vertical-align: middle;
  }

  .account-header {
    font-size: large;
    line-height: 100%;
    vertical-align: middle;

  }

  .commit {

    position: absolute;
    top: 80%;
    left: 45%;

  }

</style>
