/*! Built with http://stenciljs.com */
webcomponents.loadBundle("cntqal2x",["exports","./chunk-604eb996.js"],function(t,e){var n=window.webcomponents.h,r=function(){function t(){this.start=1,this.invokeIncrementCallback=function(){return null},this.currentlyDisplayedItems=0,this.incrementBy=10,this.initialItems=0,this.totalItems=0,this.fromLabel="from",this.moreLabel="Show more",this.buttonTheme="primary"}return t.prototype.totalItemsChanged=function(){this.reset()},t.prototype.incrementCallbackChanged=function(){this.setInvokeIncrementCallback()},t.prototype.componentWillLoad=function(){this.setInvokeIncrementCallback(),this.reset()},t.prototype.reset=function(){this.currentlyDisplayedItems=this.initialItems||this.incrementBy},t.prototype.showMore=function(){var t=this.totalItems-this.currentlyDisplayedItems;t<=0||(this.currentlyDisplayedItems+=t>this.incrementBy?this.incrementBy:t,this.invokeIncrementCallback(this.currentlyDisplayedItems))},t.prototype.setInvokeIncrementCallback=function(){this.invokeIncrementCallback=e.parseFunction(this.incrementCallback)},t.prototype.render=function(){var t=this;return n("div",null,n("span",{class:"count"},this.start," – ",this.currentlyDisplayedItems," ",this.fromLabel," ",this.totalItems),n("sdx-button",{onClick:function(){return t.showMore()},theme:this.buttonTheme},this.moreLabel))},Object.defineProperty(t,"is",{get:function(){return"sdx-show-more"},enumerable:!0,configurable:!0}),Object.defineProperty(t,"encapsulation",{get:function(){return"shadow"},enumerable:!0,configurable:!0}),Object.defineProperty(t,"properties",{get:function(){return{buttonTheme:{type:String,attr:"button-theme"},currentlyDisplayedItems:{state:!0},fromLabel:{type:String,attr:"from-label"},incrementBy:{type:Number,attr:"increment-by"},incrementCallback:{type:String,attr:"increment-callback",watchCallbacks:["incrementCallbackChanged"]},initialItems:{type:Number,attr:"initial-items"},moreLabel:{type:String,attr:"more-label"},totalItems:{type:Number,attr:"total-items",watchCallbacks:["totalItemsChanged"]}}},enumerable:!0,configurable:!0}),Object.defineProperty(t,"style",{get:function(){return":host{-webkit-box-sizing:border-box;box-sizing:border-box}*,:after,:before{-webkit-box-sizing:inherit;box-sizing:inherit}:host>div{display:-ms-flexbox;display:flex;-ms-flex-align:center;align-items:center;-ms-flex-pack:center;justify-content:center}:host>div .count{margin-right:24px}\@media (max-width:1279px){:host>div{-ms-flex-flow:column;flex-flow:column}:host>div .count{margin-bottom:8px;margin-right:0}}"},enumerable:!0,configurable:!0}),t}();t.SdxShowMore=r,Object.defineProperty(t,"__esModule",{value:!0})});