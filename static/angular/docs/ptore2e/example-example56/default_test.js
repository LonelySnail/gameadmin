describe("", function() {
  var rootEl;
  beforeEach(function() {
    rootEl = browser.rootEl;
    browser.get("build/docs/examples/example-example56/index.tpl");
  });
  
it('should check both checkBoxes', function() {
  expect(element(by.id('checkSlave')).getAttribute('checked')).toBeFalsy();
  element(by.model('master')).click();
  expect(element(by.id('checkSlave')).getAttribute('checked')).toBeTruthy();
});
});