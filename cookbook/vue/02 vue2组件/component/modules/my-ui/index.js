import MyButton from './Button'
// import MyInput from './Input'

// Vue.use(MyUI,{
//     comments:[
//       'MyButton',
//       'MyInput'
//     ]
// })

const MyUI={};

const COMMENTS=[
    MyButton // 开发的组件 export default {  name: "MyButton",
];

MyUI.install = function(Vue,options){
    console.log(options) // Vue.use(MyUI,{ --> comments:['MyButton','MyInput'] <--
    // Vue.component
    // Vue.directive
    // Vue.mixin
    // Vue.prototype.$http = function(){}

    // 判断options是存在的 options.comments也是存在的
    if(options && options.comments){
        const comments = options.comments;

        // 拿到componentName
        comments.forEach(componentName => {
           COMMENTS.forEach(component=>{
            // component(开发)     export default {  name: "MyButton",
            // componentName(使用) comments:['MyButton','MyInput']
            if(componentName===component.name){
                // 向Vue挂载组件
                Vue.component(component.name,component)
            }
           }) 
        });
    }else{
        // 直接遍历所有的组件挂载(全部加载)
        COMMENTS.forEach(component=>{
            Vue.component(component.name,component)
        })
    }
}

export default MyUI;