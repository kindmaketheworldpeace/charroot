<template>
  <div>
    <div class="talk_con" id="talk_con_id">
      <div class="chat-room-num">
         <template v-for=" i in num_arr">
            <div style="text-align: left"><span>{{i}}</span></div>

          </template>
      </div>
      <div class="content">
        <div class="talk_show" id="words">
          <div class="atalk"><span id="asay">欢迎进入聊天室</span></div>
          <template v-for=" i in message_arr">
            <div style="text-align: left"><span>{{i.Username+":" + i.Message}}</span></div>

          </template>

        </div>
        <div class="talk_input" id="talk_input_id">
          <input  v-model="message" type="text" class="talk_word" id="talkwords">
          <input @click="websocketsend()" type="button" value="发送" class="talk_sub" id="talksub">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "ChatRoom",
    data() {
      return {
        websock : null,
        message:"",
        message_arr:[],
        num_arr:[]
      }
    },
    created() {
      this.init_web_socket()
    },
    destroyed() {
      this.websock.close()

    },
    methods:{

       init_web_socket(){ //初始化weosocket
        const wsurl = "ws://127.0.0.1:9000/chatroom";
        this.websock = new WebSocket(wsurl);
        this.websock.onmessage = this.websocketonmessage;
        this.websock.onopen = this.websocketonopen;
        this.websock.onerror = this.websocketonerror;
        this.websock.onclose = this.websocketclose;
      },
      websocketonopen(){ //连接建立之后执行send方法发送数据
        // let actions = {"test":"12345"};
        // this.websocketsend(JSON.stringify(actions));
      },
      websocketonerror(){//连接建立失败重连
        this.init_web_socket();
      },
      websocketonmessage(e){ //数据接收
        var data=JSON.parse(e.data)
        console.log(data)
        if (data.Type ===1) {
          this.num_arr.push(data.Username)

        }else if (data.Type ===2) {
           this.message_arr.push(data)
        }else {

        }


      },
      websocketsend(){//数据发送
        this.websock.send(this.message);
      },
      websocketclose(e){  //关闭
        console.log('断开连接',e);
      },
    }
  }
</script>

<style scoped>
  .talk_con_mob {
    width: 600px;
    height: 500px;
    border: 1px solid #666;
    margin: 50px auto 0;
    background: #f9f9f9;
  }

  .talk_show_mob {
    width: 580px;
    height: 420px;
    border: 1px solid #666;
    background: #fff;
    margin: 10px auto 0;
    overflow: auto;
  }

  .talk_input_mob {
    width: 580px;
    margin: 10px auto 0;
  }

  .talk_word_mob {
    width: 420px;
    height: 26px;
    padding: 0px;
    float: left;
    margin-left: 10px;
    outline: none;
    text-indent: 10px;
  }

  .talk_con {
    width: 800px;
    height: 500px;
    border: 1px solid #666;
    margin: 50px auto 0;
    background: #f9f9f9;
  }

  .talk_show {
    display: inline-block;
    width: 580px;
    height: 420px;
    border: 1px solid #666;
    background: #fff;
    margin: 10px auto 0;
    overflow: auto;
  }

  .talk_input {
    display: inline-block;
    width: 580px;
    margin: 10px auto 0;
  }

  .whotalk {
    width: 80px;
    height: 30px;
    float: left;
    outline: none;
  }

  .talk_word {
    width: 490px;
    height: 26px;
    padding: 0px;
    float: left;

    outline: none;
    text-indent: 10px;
  }

  .talk_sub {
    width: 56px;
    height: 30px;
    float: right;

  }

  .atalk {
    margin: 10px;
  }

  .atalk span {
    display: inline-block;
    background: #0181cc;
    border-radius: 10px;
    color: #fff;
    padding: 5px 10px;
  }

  .btalk {
    margin: 10px;
    text-align: right;
  }

  .btalk span {
    display: inline-block;
    background: #ef8201;
    border-radius: 10px;
    color: #fff;
    padding: 5px 10px;
  }

  .chat-room-num {
    margin-left: 4px;
    display: inline-block;
    width: 20%;
    height: 460px;
    border: 1px solid #666;
    vertical-align: middle;
  }

  .content {
    display: inline-block;
    vertical-align: middle;
    width: 76%;
  }
</style>
