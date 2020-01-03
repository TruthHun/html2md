# Alert 提醒框


用于页面中操作交互后的全局提示信息，悬浮在页面顶部的中间。


- <code>Alert.info</code> 信息提醒。
- <code>Alert.success</code> 成功信息提醒。
- <code>Alert.warning</code> 警告信息提醒。
- <code>Alert.error</code> 错误信息提醒。
- <code>Alert.close</code> 关闭提醒。
- <code>Alert.closeAll</code> 关闭所有提醒。


## 获取组件




<pre><code>import { Alert } from &#39;rsuite&#39;;</code></pre>



## 演示






### 默认

<!--start-code-->




<button type="button" class="rs-btn rs-btn-default"> Info </button><button type="button" class="rs-btn rs-btn-default"> Success </button><button type="button" class="rs-btn rs-btn-default"> Warning </button><button type="button" class="rs-btn rs-btn-default"> Error </button>



<button class="rs-btn rs-btn-xs rs-btn-subtle rs-btn-icon-circle"></button>[](https://github.com/rsuite/rsuite.github.io/tree/master/src/pages/components/alert/basic.md)

<textarea>const instance = (
  &lt;ButtonToolbar&gt;
    &lt;Button onClick={() =&gt; Alert.info(&#39;This is a informations.&#39;)}&gt; Info &lt;/Button&gt;
    &lt;Button onClick={() =&gt; Alert.success(&#39;This is a successful message.&#39;)}&gt; Success &lt;/Button&gt;
    &lt;Button onClick={() =&gt; Alert.warning(&#39;This is a warning notice.&#39;)}&gt; Warning &lt;/Button&gt;
    &lt;Button onClick={() =&gt; Alert.error(&#39;This is an error message.&#39;)}&gt; Error &lt;/Button&gt;
  &lt;/ButtonToolbar&gt;
);
ReactDOM.render(instance);</textarea>

<textarea></textarea>




















<pre>xxxxxxxxxx</pre>







 





1



<pre><span role="presentation">const instance = (</pre>




2



<pre><span role="presentation">  &lt;ButtonToolbar&gt;</pre>




3



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.info(&#39;This is a informations.&#39;)}&gt; Info &lt;/Button&gt;</pre>




4



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.success(&#39;This is a successful message.&#39;)}&gt; Success &lt;/Button&gt;</pre>




5



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.warning(&#39;This is a warning notice.&#39;)}&gt; Warning &lt;/Button&gt;</pre>




6



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.error(&#39;This is an error message.&#39;)}&gt; Error &lt;/Button&gt;</pre>




7



<pre><span role="presentation">  &lt;/ButtonToolbar&gt;</pre>




8



<pre><span role="presentation">);</pre>




9



<pre><span role="presentation">ReactDOM.render(instance);</pre>

















<!--end-code-->





### 延迟关闭


duration 是一个可选项，当设置为 0 时，则不自动关闭。



<pre><code>Alert.info(content: string, duration?: number, onClose?: () =&gt; void);</code></pre>

<!--start-code-->




<button type="button" class="rs-btn rs-btn-default"> Info </button><button type="button" class="rs-btn rs-btn-default"> Success </button><button type="button" class="rs-btn rs-btn-default"> Warning </button><button type="button" class="rs-btn rs-btn-default"> Error </button>



<button class="rs-btn rs-btn-xs rs-btn-subtle rs-btn-icon-circle"></button>[](https://github.com/rsuite/rsuite.github.io/tree/master/src/pages/components/alert/duration.md)

<textarea>const instance = (
  &lt;ButtonToolbar&gt;
    &lt;Button onClick={() =&gt; Alert.info(&#39;This is a informations.&#39;, 5000)}&gt; Info &lt;/Button&gt;
    &lt;Button onClick={() =&gt; Alert.success(&#39;This is a successful message.&#39;, 5000)}&gt; Success &lt;/Button&gt;
    &lt;Button onClick={() =&gt; Alert.warning(&#39;This is a warning notice.&#39;, 5000)}&gt; Warning &lt;/Button&gt;
    &lt;Button onClick={() =&gt; Alert.error(&#39;This is an error message.&#39;, 5000)}&gt; Error &lt;/Button&gt;
  &lt;/ButtonToolbar&gt;
);
ReactDOM.render(instance);</textarea>

<textarea></textarea>




















<pre>xxxxxxxxxx</pre>







 





1



<pre><span role="presentation">const instance = (</pre>




2



<pre><span role="presentation">  &lt;ButtonToolbar&gt;</pre>




3



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.info(&#39;This is a informations.&#39;, 5000)}&gt; Info &lt;/Button&gt;</pre>




4



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.success(&#39;This is a successful message.&#39;, 5000)}&gt; Success &lt;/Button&gt;</pre>




5



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.warning(&#39;This is a warning notice.&#39;, 5000)}&gt; Warning &lt;/Button&gt;</pre>




6



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.error(&#39;This is an error message.&#39;, 5000)}&gt; Error &lt;/Button&gt;</pre>




7



<pre><span role="presentation">  &lt;/ButtonToolbar&gt;</pre>




8



<pre><span role="presentation">);</pre>




9



<pre><span role="presentation">ReactDOM.render(instance);</pre>

















<!--end-code-->





### 关闭

<!--start-code-->




<button type="button" class="rs-btn rs-btn-default">Open</button><button type="button" class="rs-btn rs-btn-default"> Close </button><button type="button" class="rs-btn rs-btn-default"> Close all </button>



<button class="rs-btn rs-btn-xs rs-btn-subtle rs-btn-icon-circle"></button>[](https://github.com/rsuite/rsuite.github.io/tree/master/src/pages/components/alert/close.md)

<textarea>const instance = (
  &lt;ButtonToolbar&gt;
    &lt;Button onClick={() =&gt; Alert.info(&#39;This is a informations.&#39;)}&gt;Open&lt;/Button&gt;
    &lt;Button onClick={() =&gt; Alert.close()}&gt; Close &lt;/Button&gt;
    &lt;Button onClick={() =&gt; Alert.closeAll()}&gt; Close all &lt;/Button&gt;
  &lt;/ButtonToolbar&gt;
);
ReactDOM.render(instance);</textarea>

<textarea></textarea>




















<pre>xxxxxxxxxx</pre>







 





1



<pre><span role="presentation">const instance = (</pre>




2



<pre><span role="presentation">  &lt;ButtonToolbar&gt;</pre>




3



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.info(&#39;This is a informations.&#39;)}&gt;Open&lt;/Button&gt;</pre>




4



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.close()}&gt; Close &lt;/Button&gt;</pre>




5



<pre><span role="presentation">    &lt;Button onClick={() =&gt; Alert.closeAll()}&gt; Close all &lt;/Button&gt;</pre>




6



<pre><span role="presentation">  &lt;/ButtonToolbar&gt;</pre>




7



<pre><span role="presentation">);</pre>




8



<pre><span role="presentation">ReactDOM.render(instance);</pre>

















<!--end-code-->




## Methods



### <code>Alert.info</code>




<pre><code>Alert.info(content: string, duration?: number, onClose?: () =&gt; void);</code></pre>



### <code>Alert.success</code>




<pre><code>Alert.success(content: string, duration?: number, onClose?: () =&gt; void);</code></pre>



### <code>Alert.warning</code>




<pre><code>Alert.warning(content: string, duration?: number, onClose?: () =&gt; void);</code></pre>



### <code>Alert.error</code>




<pre><code>Alert.error(content: string, duration?: number, onClose?: () =&gt; void);</code></pre>


**参数说明**


<table><thead><tr><th>属性名称</th><th>类型<code>(默认值)</code></th><th>描述</th></tr></thead><tbody><tr><td>content *</td><td>string</td><td>信息内容</td></tr><tr><td>duration</td><td>number <code>(2000)</code></td><td>显示的时长，超过时长后自定关闭提醒框（单位:毫秒）</td></tr><tr><td>onClose</td><td>()=&gt;void</td><td>隐藏提醒框后的回调函数</td></tr></tbody></table>



### <code>Alert.close</code>




<pre><code>Alert.close();</code></pre>



### <code>Alert.closeAll</code>




<pre><code>Alert.closeAll();</code></pre>



### <code>Alert.config</code>


全局配置



<pre><code>Alert.config(options:{
  top?: number;
  duration?: number;
  getContainer?: () =&gt; HTMLElement;
});</code></pre>



- top - 距离页面顶部的距离 (单位 px, 默认:5)
- duration - Alert 框持续时间 (默认:2000，单位: 毫秒)
- getContainer - Alert 框的父级容器


## 相关组件



- [<code>&lt;Popover&gt;</code>](./popover) 弹出框
- [<code>&lt;Tooltip&gt;</code>](./tooltip) 文字提示
- [<code>&lt;Message&gt;</code>](./message) 消息框
- [<code>&lt;Notification&gt;</code>](./notification) 通知框





[_<svg viewBox="0 0 20 20"><use xlink:href="#icon-design"></use></svg>_](/design/default/index.html#artboard1)[](https://github.com/rsuite/rsuite.github.io/edit/master/src/pages/components/alert/index.md)[](https://github.com/rsuite/rsuite/issues/new)<button type="button" class="rs-btn rs-btn-subtle language-switch-button">EN</button><button type="button" class="rs-btn rs-btn-subtle rs-btn-icon rs-btn-icon-placement-left"></button>


[获取组件](#获取组件)



[演示](#演示)

[默认](#默认)



[延迟关闭](#延迟关闭)



[关闭](#关闭)





[Methods](#Methods)

[Alert.info](#Alert.info)



[Alert.success](#Alert.success)



[Alert.warning](#Alert.warning)



[Alert.error](#Alert.error)



[Alert.close](#Alert.close)



[Alert.closeAll](#Alert.closeAll)



[Alert.config](#Alert.config)





[相关组件](#相关组件)






