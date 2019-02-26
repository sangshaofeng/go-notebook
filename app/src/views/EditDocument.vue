<template>
  <div class="main-wrapper">
    <Header></Header>
    <div class="content">
      <i-editor v-model="content" autosize></i-editor>
      <div class="btn">
          <Button @click="onSubmit" long type="success">提交</Button>
      </div>
      <article class="context" v-html="compiledMarkdown"></article>
    </div>
  </div>
</template>

<script>
import Header from '../components/Header'
import qs from 'qs'
import marked from 'marked'
export default {
  components: {
    Header
  },
  data() {
    return {
      content: '1',
    }
  },
  computed: {
    compiledMarkdown () {
    //this.articleDetail.context数据
      return marked(this.content, { sanitize: true })
    }
  },
  methods: {
    // 提交
    onSubmit () {
      console.log(this.content)
    }
  },
}
</script>

<style scoped lang="less">
.content {
  width: 1000px;
  padding: 10px 0;
  margin: 60px auto;
  text-align: left;
  background: #fff;
  box-sizing: border-box;
  textarea.ivu-input {

  }
  .btn {
    padding: 6px 11px;
    width: 100%;
  }
  .context {
    padding: 6px 11px;
    width: 100%;
  }
}
</style>