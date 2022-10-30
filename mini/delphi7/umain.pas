unit umain;

interface

uses
  Windows, Messages, SysUtils, Variants, Classes, Graphics, Controls, Forms,
  Dialogs, StdCtrls;

type
  TForm1 = class(TForm)
    Button1: TButton;
    Edit1: TEdit;
    procedure Button1Click(Sender: TObject);
  private
    { Private declarations }
  public
    { Public declarations }
  end;

//libgolang dl ÖÐµÄº¯Êý
//function DefDlgProc(hDlg: HWND; Msg: UINT; wParam: WPARAM; lParam: LPARAM): LRESULT; cdecl;  external gdi32 name 'IntersectClipRect';
function run_golang_json(src_json:PAnsiChar): PAnsiChar; cdecl;  external 'libgolang.dll';
procedure free_golang_cstr(golang_c_str:PAnsiChar); cdecl;  external 'libgolang.dll';
//export free_golang_cstr
//func free_golang_cstr(golang_c_str *C.char)  {


//
//void (*func_on_event)(const char * event, const char * key, const char * value);
procedure func_on_event(const event:PAnsiChar; const key:PAnsiChar; const value:PAnsiChar); cdecl;

var
  Form1: TForm1;

implementation

{$R *.dfm}

procedure TForm1.Button1Click(Sender: TObject);
var
  s:string;
  ps:PAnsiChar;
begin
  ps := run_golang_json('');


  s := StrPas(ps);

  free_golang_cstr(ps);

  //----

  ps := run_golang_json(PAnsiChar('{"func_name":"set_on_event", "param1":"' + inttostr(Integer(@func_on_event)) + '"}'));


  s := StrPas(ps);

  free_golang_cstr(ps);

end;


procedure func_on_event(const event:PAnsiChar; const key:PAnsiChar; const value:PAnsiChar); cdecl;
begin
  //
  ShowMessage(event);


end;  

end.
