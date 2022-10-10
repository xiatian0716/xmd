<template>
  <button :class="['my-btn',type]">
    <slot></slot>
  </button>
</template>

<script>
export default {
  name: "MyButton",
  props:{
    // 配置css属性
    type:String
  }
}
</script>

<style lang="scss">
.my-btn{
    height: 34px;
    // 移动到了
    // common.css
    // border: none;
    // outline: none;
    padding:0 15px;
    border:1px solid #ddd;
    background-color: #fff;

    &.primary{
      border:1px solid #ddd;
      background-color: blue;
      color:#fff;
    }

    &.danger{
      border:1px solid #ddd;
      background-color: red;
      color:#fff;
    }
}
</style>